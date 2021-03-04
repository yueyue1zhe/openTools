<?php
/**
 * util.php
 * Created by PhpStorm
 * Author: zhouzhe5934@icloud.com
 * Date  : 2021/1/26
 */

const ApiEntryKey = "yroute";
const ApiScopeFrontend = "frontend";

IModulesLoader::AppStart();
class IModulesLoader{
    public static function AppStart(){
        spl_autoload_register('IModulesLoader::autoload');
    }
    public static $vendorMap = [
        'service' => __DIR__ . DIRECTORY_SEPARATOR . 'service',
        'model' => __DIR__ . DIRECTORY_SEPARATOR . 'model',
        ApiScopeFrontend => __DIR__ . DIRECTORY_SEPARATOR . ApiScopeFrontend,
    ];
    public static function autoload($class){
        $file = self::findFile($class);
        if (file_exists($file)) {
            self::includeFile($file);
        }
    }
    public static function findFile($class){
        $vendor = explode("\\",$class);
        $vendorDir = self::$vendorMap[$vendor[1]];
        $filePath = substr($class, strlen($vendor[0].DIRECTORY_SEPARATOR.$vendor[1])) . '.php';
        return strtr($vendorDir . $filePath, '\\', DIRECTORY_SEPARATOR);
    }
    private static function includeFile($file)    {
        if (is_file($file)) {
            include $file;
        }
    }
}

class BaseController{

    protected $token;
    protected $api_scope;
    protected $api_judge_si = true;
    protected $auth_route = true;
    protected $post;

    public function __construct() {
        global $_GPC;
        $this->post = $_GPC["__input"];
        if (!$this->auth_route)return;
        $token = AppUtil::visitToken();
        if (empty($token) || is_error($token))AppUtil::ReqLoginFail($token["message"] ?: "token无效");
        if ($this->api_judge_si)$this->token_judge_si($token);
        $this->token = $token;
    }

    protected function token_award($data, $jwtUuid){
        return \Jwt::awardToken($data,$this->api_scope,$jwtUuid);
    }
    protected function token_judge_j(){if (empty($this->token))return;}
    private function token_judge_si($token){
        if (empty($token)){
            AppUtil::ReqLoginFail("token无效");
        }
        $origin = isset($_SERVER['HTTP_ORIGIN'])? $_SERVER['HTTP_ORIGIN'] : '';
        if (!empty($token["sub"]) && $origin != $token["sub"]){
            AppUtil::ReqLoginFail("非授权域");
        }
        if ($token["iss"] != $this->api_scope){
            AppUtil::ReqLoginFail("权限异常");
        }
    }
}

class AppUtil{
    public static function MakeOrderNo($pre=""){
        $yCode = array('A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J');
        return $pre.$yCode[intval(date('Y')) - 2021] . strtoupper(dechex(date('m'))) . date('d') . substr(time(), -5) . substr(microtime(), 2, 5) . sprintf('%02d', rand(0, 99));
    }
    public static function visitToken(){
        return Jwt::verifyToken($_SERVER['HTTP_TOKEN']);
    }
    public static function getMarRF($scope){
        global $_GPC;
        $space = current(explode("_",ModuleName));
        $use = explode("/",ltrim($_GPC[ApiEntryKey],"/"));
        $func = array_pop($use);
        $use = $use ? "\\".implode("\\",$use) : "";
        return [
            "route" =>$space."\\".$scope.$use,
            "func" => $func
        ];
    }
    private static function setMarCors($allowOrigin){
        if ($_SERVER['REQUEST_METHOD'] == 'GET')exit();
        if (!empty($allowOrigin)){
            $origin = isset($_SERVER['HTTP_ORIGIN'])? $_SERVER['HTTP_ORIGIN'] : '';
            if (in_array($origin,$allowOrigin)){
                header("Content-Type: application/json");
                header("access-control-allow-headers: token,content-type");
                header("access-control-allow-methods: POST");
                header("access-control-allow-origin: {$origin}");
            }
        }
        if ('POST' != $_SERVER['REQUEST_METHOD'])exit();
        if (!isset($_SERVER['REQUEST_METHOD']))exit();
    }
    public static function Mar($ApiScope,$allowOrigin=[]){
        self::setMarCors($allowOrigin);
        $cf = self::getMarRF($ApiScope);
        if(!class_exists($cf["route"]))self::ReqFail("内部错误:路由不存在");
        $apiClass = new $cf["route"]($cf);
        if(!method_exists($apiClass, $cf["func"]))self::ReqFail("内部错误:方法不存在");
        $result = call_user_func_array([$apiClass, $cf["func"]], []);
        if(is_error($result))self::ReqFail($result["message"],$result["errno"]);
        self::ReqOk($result);
    }

    public static function ReqLoginFail($message,$errno=40019){
        self::Req("",$message,$errno);
    }
    public static function ReqOk($data){
        self::Req($data,"",0);
    }
    public static function ReqFail($message,$errno=1){
        self::Req("",$message,$errno);
    }
    public static function Req($data,$message,$errno){
        header("Content-Type: application/json; charset=utf-8");
        echo json_encode(["data" => $data,"message" => $message,"errno" => $errno]);
        exit();
    }
    public static function UrlIsImage($url){
        return !empty($url) && stripos(get_headers($url)[1],"image") !== false;
    }
    public static function FillFieldStruct($data,$struct){
        $new = [];
        foreach ($data as $k => $v){if (in_array($k,$struct))$new[$k] = $v;}
        return $new;
    }
    public static function Emoji2Str($value,$re="*"){
        $value = json_encode($value);
        $value = preg_replace("/\\\u[ed][0-9a-f]{3}\\\u[ed][0-9a-f]{3}/",$re,$value);
        return json_decode($value);
    }
    public static function ArrKeys($arr, $key="id"){
        $tmp = array();
        if(!empty($arr)){
            foreach ($arr as $k => $v){
                array_push($tmp,$v[$key]);
            }
        }
        return $tmp;
    }
    public static function MakeTodayTime($h,$m,$s){
        return mktime($h,$m,$s,date('m'),date('d'),date('Y'));
    }
}

class Jwt {

    //头部
    private static $header=array(
        'alg'=>'HS256', //生成signature的算法
        'typ'=>'JWT'    //类型
    );

    //使用HMAC生成信息摘要时所使用的密钥
    private static $key='123456BBACDDFzasedCC';

    public static function jwtUuid(){
        return md5(uniqid('JWT').time());
    }
    public static function awardToken($data,$iss,$jwtUuid="",$exp=7200,$sub=""){
        $payload = array(
            'data'=> $data,
            'iss'=>$iss,
            'iat'=>time(),
            'exp'=>time()+$exp,
            'nbf'=>time(),
            'sub'=> $sub ? $sub : isset($_SERVER['HTTP_ORIGIN'])? $_SERVER['HTTP_ORIGIN'] : '',
            'jti'=> $jwtUuid ? $jwtUuid : self::jwtUuid()
        );
        return \Jwt::getToken($payload);
    }

    /**
     * 获取jwt token
     * @param array $payload jwt载荷   格式如下非必须
     * [
     *  'iss'=>'jwt_admin',  //该JWT的签发者
     *  'iat'=>time(),  //签发时间
     *  'exp'=>time()+7200,  //过期时间
     *  'nbf'=>time()+60,  //该时间之前不接收处理该Token
     *  'sub'=>'www.admin.com',  //面向的用户
     *  'jti'=>md5(uniqid('JWT').time())  //该Token唯一标识
     * ]
     * @return bool|string
     */
    public static function getToken($payload)
    {
        if(is_array($payload))
        {
            $base64header=self::base64UrlEncode(json_encode(self::$header,JSON_UNESCAPED_UNICODE));
            $base64payload=self::base64UrlEncode(json_encode($payload,JSON_UNESCAPED_UNICODE));
            $token=$base64header.'.'.$base64payload.'.'.self::signature($base64header.'.'.$base64payload,self::$key,self::$header['alg']);
            return $token;
        }else{
            return false;
        }
    }


    /**
     * 验证token是否有效,默认验证exp,nbf,iat时间
     * @param string $Token 需要验证的token
     * @return bool|string
     */
    public static function verifyToken($Token)
    {
        $tokens = explode('.', $Token);
        if (count($tokens) != 3)
            return false;

        list($base64header, $base64payload, $sign) = $tokens;

        //获取jwt算法
        $base64decodeheader = json_decode(self::base64UrlDecode($base64header), JSON_OBJECT_AS_ARRAY);
        if (empty($base64decodeheader['alg']))
            return false;

        //签名验证
        if (self::signature($base64header . '.' . $base64payload, self::$key, $base64decodeheader['alg']) !== $sign)
            return false;

        $payload = json_decode(self::base64UrlDecode($base64payload), JSON_OBJECT_AS_ARRAY);

        //签发时间大于当前服务器时间验证失败
        if (isset($payload['iat']) && $payload['iat'] > time())
            return false;

        //过期时间小宇当前服务器时间验证失败
        if (isset($payload['exp']) && $payload['exp'] < time())
            return false;

        //该nbf时间之前不接收处理该Token
        if (isset($payload['nbf']) && $payload['nbf'] > time())
            return false;

        return $payload;
    }




    /**
     * base64UrlEncode   https://jwt.io/  中base64UrlEncode编码实现
     * @param string $input 需要编码的字符串
     * @return string
     */
    private static function base64UrlEncode($input)
    {
        return str_replace('=', '', strtr(base64_encode($input), '+/', '-_'));
    }

    /**
     * base64UrlEncode  https://jwt.io/  中base64UrlEncode解码实现
     * @param string $input 需要解码的字符串
     * @return bool|string
     */
    private static function base64UrlDecode($input)
    {
        $remainder = strlen($input) % 4;
        if ($remainder) {
            $addlen = 4 - $remainder;
            $input .= str_repeat('=', $addlen);
        }
        return base64_decode(strtr($input, '-_', '+/'));
    }

    /**
     * HMACSHA256签名   https://jwt.io/  中HMACSHA256签名实现
     * @param string $input 为base64UrlEncode(header).".".base64UrlEncode(payload)
     * @param string $key
     * @param string $alg   算法方式
     * @return mixed
     */
    private static function signature($input, $key, $alg = 'HS256')
    {
        $alg_config=array(
            'HS256'=>'sha256'
        );
        return self::base64UrlEncode(hash_hmac($alg_config[$alg], $input, $key,true));
    }
}

class W7DBBase extends We7Table {
    public function __construct() {
        parent::__construct();
    }
    public function Uniacid(){
        global $_W;
        $this->query->where("uniacid",$_W["uniacid"]);
        return $this;
    }
    public function ByIdWithUniacid($id){
        global $_W;
        return $this->getById($id,$_W["uniacid"]);
    }
    public function FillTableField($param,$updatetime=false){
        if ($updatetime)$param["updatetime"] = TIMESTAMP;
        return AppUtil::FillFieldStruct($param,$this->field);
    }

    public function WhereId($id){
        $this->query->where("id",$id);
        return $this;
    }

    public function Total($con=[],$withUniacid=true){
        global $_W;
        if ($withUniacid && !in_array("uniacid",$con["uniacid"]))$con["uniacid"]=$_W["uniacid"];
        return pdo_count($this->getTableName(),$con,0);
    }

    public function UniacidGetPaging($con,$page,$size=15,$fields=[],$orderBy="createtime Desc"){
        global $_W;
        $row = array();
        $page = max(1, intval($page));
        if (empty($con["uniacid"]))$con["uniacid"] = $_W["uniacid"];
        $row["total"] = $this->Total($con);
        $row["total"] = intval($this->UniacidGetOne($con, array("COUNT(*) as total"))["total"]);
        $row["list"] = $this->UniacidGetAll($con, $fields, $orderBy,$page, $size);
        $row["page"] = $page;
        $row["size"] = $size;
        return $row;
    }
    public function UniacidGetAll($con = array(), $fields=array(),$orderBy = "createtime DESC", $offset=0, $size=0) {
        global $_W;
        $size ? $limit = [$offset,$size] : $limit = "";
        if (empty($con["uniacid"]))$con["uniacid"] = $_W["uniacid"];
        return pdo_getall($this->getTableName(), $con, $fields, '', $orderBy, $limit);
    }
    public function UniacidGetOne($con=array(), $fields = array()){
        global $_W;
        if (empty($con["uniacid"]))$con["uniacid"] = $_W["uniacid"];
        return pdo_get($this->getTableName(),$con, $fields);
    }
    public function UniacidUpdate($data, $con=[],$updatetime=true):int{
        global $_W;
        if (empty($con["uniacid"]))$con["uniacid"] = $_W["uniacid"];
        if ($updatetime)$data["updatetime"] = TIMESTAMP;
        return pdo_update($this->getTableName(),$data,$con);
    }
    public function UniacidDel($con){
        if (empty($con))return;
        global $_W;
        if (empty($con["uniacid"]))$con["uniacid"] = $_W["uniacid"];
        pdo_delete($this->getTableName(),$con);
    }
}

class W7Util{

    public static function ExcelExport($list,$filename,$indexKey,$startRow=1,$excel2007=true){
        load()->library("phpexcel/PHPExcel");
        if(empty($filename)) $filename = time();
        if( !is_array($indexKey)) return;
        $header_arr = array('A','B','C','D','E','F','G','H','I','J','K','L','M', 'N','O','P','Q','R','S','T','U','V','W','X','Y','Z');
        $objPHPExcel = new \PHPExcel();
        if($excel2007){
            $objWriter = new PHPExcel_Writer_Excel2007($objPHPExcel);
            $filename = $filename.'.xlsx';
        }else{
            $objWriter = new PHPExcel_Writer_Excel5($objPHPExcel);
            $filename = $filename.'.xls';
        }
        $objActSheet = $objPHPExcel->getActiveSheet();
        foreach ($list as $row) {
            foreach ($indexKey as $key => $value){
                $objActSheet->setCellValue($header_arr[$key].$startRow,$row[$value]);
            }
            $startRow++;
        }
        header("Pragma: public");
        header("Expires: 0");
        header("Cache-Control:must-revalidate, post-check=0, pre-check=0");
        header("Content-Type:application/force-download");
        header("Content-Type:application/vnd.ms-execl");
        header("Content-Type:application/octet-stream");
        header("Content-Type:application/download");;
        header('Content-Disposition:attachment;filename='.$filename.'');
        header("Content-Transfer-Encoding:binary");
        $objWriter->save('php://output');
    }

    /**
     * 二次封装插件函数式调用 移动端
     * @param string $module_name 模块名
     * @param string $func 方法名
     * @param array $param 参数
     * @return mixed
     */
    public static function MobileHookGet($module_name,$func,$param=array()){
        return self::HookGet($module_name,"hookMobile".$func,$param);
    }

    /**
     * 二次封装插件函数式调用 pc端
     * @param string $module_name 模块名
     * @param string $func 方法名
     * @param array $param 参数
     * @return mixed
     */
    public static function WebHookGet($module_name,$func,$param=array()){
        return self::HookGet($module_name,"hookWeb".$func,$param);
    }

    /**
     * 微擎插件函数式调用
     * @param $module_name
     * @param $func
     * @param array $params
     * @return mixed
     */
    public static function HookGet($module_name,$func,$params=array()) {
        load()->model('module');
        $plugin_info = module_fetch($module_name);
        if (empty($plugin_info)) {
            return error(1,"{$module_name}插件不存在");
        }
        $plugin_module = WeUtility::createModuleHook($module_name);
        if (method_exists($plugin_module, $func) && $plugin_module instanceof WeModuleHook) {
            return call_user_func_array(array($plugin_module, $func), array('params' => $params));
        }else{
            return error(1,"模块 {$module_name} 不存在嵌入点 {$func}");
        }
    }

    public static function DoseMedia($url){
        global $_W;
        if(empty($url))return $url;
        if(stripos($url,"http") !== 0)$url = $_W["attachurl"].$url;
        return $url;
    }
    public static function FormatMedia($url){
        global $_W;
        if(empty($url))return $url;
        if(stripos($url,$_W["attachurl"]) === 0)$url = str_replace($_W["attachurl"],"",$url);
        return $url;
    }
    public static function MakeQrCode2Show($content){
        load()->library("qrcode");
        QRcode::png($content,false,QR_ECLEVEL_H,12,1);
    }

    public static function CheckKeyWord($content,$type,$status=1,$displayOrder=1,$moduleName=ModuleName){
        global $_W;
        return pdo_get("rule_keyword",array(
            "uniacid" => $_W["uniacid"],
            "module" => $moduleName,
            "content" => $content,
            "type" => $type,
            "displayorder" =>$displayOrder,
            "status" =>$status
        ));
    }

    /**
     * 微擎关键字清理
     * @param string $keyword 关键字内容
     * @param string $module_name 模块名称
     */
    public static function ClearKeyWord($keyword,$module_name=ModuleName){
        global $_W;
        $key = pdo_get("rule_keyword",array(
            "uniacid" => $_W["uniacid"],
            "module" => $module_name,
            "content" => $keyword
        ));
        if (!empty($key)){
            pdo_delete("rule_keyword",["id"=>$key["id"]]);
            $rule = pdo_get("rule",[
                'id' => $key["rid"],
                'uniacid' => $_W['uniacid'],
                'module' => $module_name
            ]);
            if(!empty($rule)){
                pdo_delete("rule",["id"=>$rule["id"]]);
            }
        }
    }
    /**
     * 微擎关键字创建
     * @param string $keyword 关键字文本
     * @param string $module_name 模块名称
     * @param int $type 1精准2包含3正则
     * @param int $status 1开启 0关闭
     * @param int $displayOrder 优先级
     */
    public static function FetchKeyWord($keyword,$type=1,$status=1,$displayOrder=1,$module_name=ModuleName){
        global $_W;
        $key = pdo_get("rule_keyword",array(
            "uniacid" => $_W["uniacid"],
            "module" => $module_name,
            "content" => $keyword,
            "type" => $type,
            "displayorder" =>$displayOrder,
            "status" =>$status
        ));
        if (empty($key)){
            $rule = pdo_get("rule",['name'=>$keyword,'uniacid' => $_W['uniacid'],'module' => $module_name]);
            if(!empty($rule)){
                $rid = $rule["id"];
            }else{
                pdo_insert("rule",[
                    'uniacid' => $_W["uniacid"],
                    'name' => $keyword,
                    'module' => $module_name,
                    'displayorder'=>$displayOrder,
                    'status' => $status
                ]);
                $rid = pdo_insertid();
            }
            pdo_insert("rule_keyword",array(
                "rid" => $rid,
                "uniacid" => $_W["uniacid"],
                "module" => $module_name,
                "content" => $keyword,
                "type" => $type,
                "displayorder" =>$displayOrder,
                "status" =>$status
            ));
        }
    }

    /**
     * 搜索公众号粉丝
     * @param $value
     * @return mixed
     */
    public static function SearchFansFunc($value){
        $value = trim($value);
        if (empty($value))return error(1,"搜索条件不能为空");
        $fans = pdo_getall("mc_mapping_fans",["uid >"=>0,"follow"=>1,"nickname LIKE"=>"%{$value}%"],["uid","openid","nickname"],"","followtime DESC");
        if (empty($fans))return error(1,"搜索结果为空");
        foreach ($fans as $key => $value){
            $tmp = mc_fansinfo($value["openid"]);
            $fans[$key]["avatar"] = $tmp["tag"]["avatar"];
        }
        return $fans;
    }
    public static function SearchFans($name="",$value="",$tap="searchFans",$url=""){
        return "
<section style='display: flex;width: 100%'>
    <input onchange='`{$name}change()`' type=\"text\" id='set{$name}' name=\"{$name}\" class=\"form-control\" value=\"{$value}\" />
    <span class=\"btn btn-default\" data-toggle=\"modal\" data-target=\"#{$name}\">选择粉丝</span>
</section>
<div class=\"modal fade\" id=\"{$name}\" tabindex=\"-1\" role=\"dialog\" aria-labelledby=\"myModalLabel\" aria-hidden=\"true\">
    <div class=\"modal-dialog\">
        <div class=\"modal-content\">
            <div class=\"modal-header\">
                <button id='hide{$name}' type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-hidden=\"true\">&times;</button>
                <h4 class=\"modal-title\" id=\"myModalLabel\">选择粉丝</h4>
            </div>
            <div class=\"modal-body\">
                <div class=\"form-group\">
                    <label class=\"col-sm-2 control-label\" style=\"text-align:left;\">粉丝昵称</label>
                    <div style='display: flex' class=\"col-sm-8\">
                        <input id='{$name}search' type=\"text\" class=\"form-control\" value=\"\" />
                        <span onclick='{$name}searchFans()' class='btn btn-default'>搜索</span>
                    </div>
                </div>
                <section id='{$name}fansListBox' style='height: 35rem;overflow-x: auto'>
                                 
                </section>
            </div>
            <div class=\"modal-footer\"></div>
        </div>
    </div>
</div>
<script>
    function {$name}chooseThisFansOpenid(openid) {
        document.getElementById(`set{$name}`).value = openid;
        document.getElementById(`hide{$name}`).click();
    }
    function {$name}searchFans() {
        let search = document.getElementById(`{$name}search`).value;
        $.post(`{$url}`,{value: search,tap:`{$tap}`},res=>{
            if (res.errno != 0){
                alert(res.message);
                return;
            }
            $('#{$name}fansListBox').empty();
            let str = ``;
            for (let key in res.data){
                if (res.data.hasOwnProperty(key)){
                    str += `<section style='width: 100%; display: flex;border-top: 2px solid #ededed;padding: 0.5rem'>
                        <section style='width: 33%;margin: auto'><img style='width: 4rem;height: 4rem' src='`+res.data[key].avatar+`' alt=''></section>
                        <span style='width: 33%;margin: auto'>`+res.data[key].nickname+`</span>
                        <section style='width: 33%;margin: auto;text-align: right'><span onclick='{$name}chooseThisFansOpenid(\"`+res.data[key].openid+`\")' class='btn btn-default'>选择</span></section>
                    </section>`;
                }
            }
            $('#{$name}fansListBox').append(str);
        })
    }
</script>
";
    }
    public static function SearchFansSubmit($tap="searchFans"){
        global $_W,$_GPC;
        if (!$_W["isajax"] || $_GPC["tap"] != $tap)return null;
        $row =  self::SearchFansFunc($_GPC["value"]);
        if (is_error($row))AppUtil::ReqFail($row["message"]);
        AppUtil::ReqOk($row);
    }

    public static function CoreModuleSitePay($fee,$tid,$title){
        $_POST["module"] = ModuleName;
        $_POST["method"] = "wechat";
        $_POST["fee"] = $fee;
        $_POST["tid"] = $tid;
        $_POST["title"] = $title;
        $site = WeUtility::createModuleSite("core");
        $site->doMobilePay();
    }

    public static function InitGlobal($openid){
        global $_W;
        $_W['openid'] = $openid;
        $_W['fans'] = mc_fansinfo($_W['openid']);
        $_W['fans']['from_user'] = $_W['fans']['openid'] = $_W['openid'];
        $_W['member'] = mc_fetch($_W['openid']);
    }

    public static function GetWxJsSdkCfg($url=""){
        $account_api = WeAccount::create();
        return $account_api->getJssdkConfig($url);
    }

    public static function WxCode2FansInfo($code){
        $oauth_account = \WeAccount::createByUniacid();
        $oauth = $oauth_account->getOauthInfo($code);
        if (is_error($oauth)){
            return $oauth;
        }
        self::wxCode2FansInfoSync($oauth_account,$oauth);
        return mc_fansinfo($oauth["openid"]);
    }
    private static function wxCode2FansInfoSync($oauth_account,$oauth){
        global $_W;
        $scope = $oauth["scope"];
        if (intval($_W['account']['level']) == ACCOUNT_SERVICE_VERIFY) {
            $fan = mc_fansinfo($oauth['openid']);
            if (!empty($fan)) {
                $_SESSION['openid'] = $oauth['openid'];
                if (empty($_SESSION['uid'])) {
                    if (!empty($fan['uid'])) {
                        $member = mc_fetch($fan['uid'], array('uid'));
                        if (!empty($member) && $member['uniacid'] == $_W['uniacid']) {
                            $_SESSION['uid'] = $member['uid'];
                        }
                    }
                }
            } else {
                $accObj = WeAccount::createByUniacid($_W['uniacid']);
                $userinfo = $accObj->fansQueryInfo($oauth['openid']);

                if(!is_error($userinfo) && !empty($userinfo) && !empty($userinfo['subscribe'])) {
                    $userinfo['nickname'] = stripcslashes($userinfo['nickname']);
                    $userinfo['avatar'] = $userinfo['headimgurl'];
                    $_SESSION['userinfo'] = base64_encode(iserializer($userinfo));
                    $record = array(
                        'openid' => $userinfo['openid'],
                        'uid' => 0,
                        'acid' => $_W['acid'],
                        'uniacid' => $_W['uniacid'],
                        'salt' => random(8),
                        'updatetime' => TIMESTAMP,
                        'nickname' => stripslashes($userinfo['nickname']),
                        'follow' => $userinfo['subscribe'],
                        'followtime' => $userinfo['subscribe_time'],
                        'unfollowtime' => 0,
                        'unionid' => $userinfo['unionid'],
                        'tag' => base64_encode(iserializer($userinfo)),
                        'user_from' => $_W['account']->typeSign == 'wxapp' ? 1 : 0,
                    );

                    if (!isset($unisetting['passport']) || empty($unisetting['passport']['focusreg'])) {
                        $email = md5($oauth['openid']).'@we7.cc';
                        $email_exists_member = table('mc_members')
                            ->where(array(
                                'email' => $email,
                                'uniacid' => $_W['uniacid']
                            ))
                            ->getcolumn('uid');
                        if (!empty($email_exists_member)) {
                            $uid = $email_exists_member;
                        } else {
                            $default_groupid = table('mc_groups')
                                ->where(array(
                                    'uniacid' => $_W['uniacid'],
                                    'isdefault' => 1
                                ))
                                ->getcolumn('groupid');
                            $data = array(
                                'uniacid' => $_W['uniacid'],
                                'email' => $email,
                                'salt' => random(8),
                                'groupid' => $default_groupid,
                                'createtime' => TIMESTAMP,
                                'password' => md5($message['from'] . $data['salt'] . $_W['config']['setting']['authkey']),
                                'nickname' => stripslashes($userinfo['nickname']),
                                'avatar' => $userinfo['headimgurl'],
                                'gender' => $userinfo['sex'],
                                'nationality' => $userinfo['country'],
                                'resideprovince' => $userinfo['province'] . '省',
                                'residecity' => $userinfo['city'] . '市',
                            );
                            table('mc_members')->fill($data)->save();
                            $uid = pdo_insertid();
                        }
                        $record['uid'] = $uid;
                        $_SESSION['uid'] = $uid;
                    }
                    table('mc_mapping_fans')->fill($record)->save();
                    $mc_fans_tag_table = table('mc_fans_tag');
                    $mc_fans_tag_fields = mc_fans_tag_fields();
                    $fans_tag_update_info = array();
                    foreach ($userinfo as $fans_field_key => $fans_field_info) {
                        if (in_array($fans_field_key, array_keys($mc_fans_tag_fields))) {
                            $fans_tag_update_info[$fans_field_key] = $fans_field_info;
                        }
                        $fans_tag_update_info['tagid_list'] = iserializer($fans_tag_update_info['tagis_list']);
                    }
                    $fans_tag_exists = $mc_fans_tag_table->getByOpenid($fans_tag_update_info['openid']);
                    if (!empty($fans_tag_exists)) {
                        table('mc_fans_tag')
                            ->where(array('openid' => $fans_tag_update_info['openid']))
                            ->fill($fans_tag_update_info)
                            ->save();
                    } else {
                        table('mc_fans_tag')->fill($fans_tag_update_info)->save();
                    }
                } else {
                    $record = array(
                        'openid' => $oauth['openid'],
                        'nickname' => '',
                        'subscribe' => '0',
                        'subscribe_time' => '',
                        'headimgurl' => '',
                    );
                }
                $_SESSION['openid'] = $oauth['openid'];
                $_W['fans'] = $record;
                $_W['fans']['from_user'] = $record['openid'];
            }
        }
        if (intval($_W['account']['level']) != ACCOUNT_SERVICE_VERIFY) {
            $mc_oauth_fan = mc_oauth_fans($oauth['openid'], $_W['uniacid']);
            if (empty($mc_oauth_fan)) {
                $data = array(
                    'uniacid' => $_W['uniacid'],
                    'oauth_openid' => $oauth['openid'],
                    'uid' => intval($_SESSION['uid']),
                    'openid' => $_SESSION['openid']
                );
                table('mc_oauth_fans')->fill($data)->save();
            }
            //如果包含Unionid，则直接查原始openid
            if (!empty($oauth['unionid'])) {
                $fan = table('mc_mapping_fans')
                    ->searchWithUnionid($oauth['unionid'])
                    ->searchWithUniacid($_W['uniacid'])
                    ->get();
                if (!empty($fan)) {
                    if (!empty($fan['uid'])) {
                        $_SESSION['uid'] = intval($fan['uid']);
                    }
                    if (!empty($fan['openid'])) {
                        $_SESSION['openid'] = strval($fan['openid']);
                    }
                }
            } else {
                if (!empty($mc_oauth_fan)) {
                    if (empty($_SESSION['uid']) && !empty($mc_oauth_fan['uid'])) {
                        $_SESSION['uid'] = intval($mc_oauth_fan['uid']);
                    }
                    if (empty($_SESSION['openid']) && !empty($mc_oauth_fan['openid'])) {
                        $_SESSION['openid'] = strval($mc_oauth_fan['openid']);
                    }
                }
            }
        }
        if ($scope == 'userinfo' || $scope == 'snsapi_userinfo') {
            $userinfo = $oauth_account->getOauthUserInfo($oauth['access_token'], $oauth['openid']);
            if (!is_error($userinfo)) {
                $userinfo['nickname'] = stripcslashes($userinfo['nickname']);
                $userinfo['avatar'] = $userinfo['headimgurl'];
                $_SESSION['userinfo'] = base64_encode(iserializer($userinfo));
                $fan = table('mc_mapping_fans')->searchWithOpenid($oauth['openid'])->searchWithUniacid($_W['uniacid'])->get();
                if (!empty($fan)) {
                    $record = array();
                    $record['updatetime'] = TIMESTAMP;
                    $record['nickname'] = stripslashes($userinfo['nickname']);
                    $record['tag'] = base64_encode(iserializer($userinfo));
                    if (empty($fan['unionid'])) {
                        $record['unionid'] = !empty($userinfo['unionid']) ? $userinfo['unionid'] : '';
                    }
                    table('mc_mapping_fans')
                        ->where(array(
                            'openid' => $fan['openid'],
                            'uniacid' => $_W['uniacid']
                        ))
                        ->fill($record)
                        ->save();
                    if (!empty($fan['uid']) || !empty($_SESSION['uid'])) {
                        $uid = $fan['uid'];
                        if(empty($uid)){
                            $uid = $_SESSION['uid'];
                        }
                        $user = mc_fetch($uid, array('nickname', 'gender', 'residecity', 'resideprovince', 'nationality', 'avatar'));
                        $record = array();
                        if(empty($user['nickname']) && !empty($userinfo['nickname'])) {
                            $record['nickname'] = stripslashes($userinfo['nickname']);
                        }
                        if(empty($user['gender']) && !empty($userinfo['sex'])) {
                            $record['gender'] = $userinfo['sex'];
                        }
                        if(empty($user['residecity']) && !empty($userinfo['city'])) {
                            $record['residecity'] = $userinfo['city'] . '市';
                        }
                        if(empty($user['resideprovince']) && !empty($userinfo['province'])) {
                            $record['resideprovince'] = $userinfo['province'] . '省';
                        }
                        if(empty($user['nationality']) && !empty($userinfo['country'])) {
                            $record['nationality'] = $userinfo['country'];
                        }
                        if(empty($user['avatar']) && !empty($userinfo['headimgurl'])) {
                            $record['avatar'] = $userinfo['headimgurl'];
                        }
                        if(!empty($record)) {
                            mc_update($user['uid'], $record);
                        }
                    }
                } else {
                    $record = array(
                        'openid' => $oauth['openid'],
                        'uid' => 0,
                        'acid' => $_W['acid'],
                        'uniacid' => $_W['uniacid'],
                        'salt' => random(8),
                        'updatetime' => TIMESTAMP,
                        'nickname' => $userinfo['nickname'],
                        'follow' => 0,
                        'followtime' => 0,
                        'unfollowtime' => 0,
                        'tag' => base64_encode(iserializer($userinfo)),
                        'unionid' => !empty($userinfo['unionid']) ? $userinfo['unionid'] : '',
                        'user_from' => $_W['account']->typeSign == 'wxapp' ? 1 : 0,
                    );

                    if (!isset($unisetting['passport']) || empty($unisetting['passport']['focusreg'])) {
                        $default_groupid = table('mc_groups')
                            ->where(array(
                                'uniacid' => $_W['uniacid'],
                                'isdefault' => 1
                            ))
                            ->getcolumn('groupid');
                        $data = array(
                            'uniacid' => $_W['uniacid'],
                            'email' => md5($oauth['openid']).'@we7.cc',
                            'salt' => random(8),
                            'groupid' => $default_groupid,
                            'createtime' => TIMESTAMP,
                            'password' => md5($message['from'] . $data['salt'] . $_W['config']['setting']['authkey']),
                            'nickname' => $userinfo['nickname'],
                            'avatar' => $userinfo['headimgurl'],
                            'gender' => $userinfo['sex'],
                            'nationality' => $userinfo['country'],
                            'resideprovince' => $userinfo['province'] . '省',
                            'residecity' => $userinfo['city'] . '市',
                        );
                        table('mc_members')
                            ->fill($data)
                            ->save();
                        $uid = pdo_insertid();
                        $record['uid'] = $uid;
                        $_SESSION['uid'] = $uid;
                    }
                    table('mc_mapping_fans')->fill($record)->save();
                }
            } else {
                //                message('微信授权获取用户信息失败,错误信息为: ' . $response['message']);
            }
        }
    }

    public static function FastModulePathBase($path){
        return IA_ROOT.DIRECTORY_SEPARATOR."addons".DIRECTORY_SEPARATOR.ModuleName.$path;
    }
    public static function uploadRemoteFile($filePath,$filename,$auto_delete_local=true){
        global $_W;
        if (empty($_W['setting']['remote']['type'])) {
            return false;
        }
        if ($_W['setting']['remote']['type'] == ATTACH_FTP) {
            load()->library('ftp');
            $ftp_config = array(
                'hostname' => $_W['setting']['remote']['ftp']['hostname'] ?: $_W['setting']['remote']['ftp']['host'],
                'username' => $_W['setting']['remote']['ftp']['username'],
                'password' => $_W['setting']['remote']['ftp']['password'],
                'port' => $_W['setting']['remote']['ftp']['port'],
                'ssl' => $_W['setting']['remote']['ftp']['ssl'],
                'passive' => $_W['setting']['remote']['ftp']['passive'] ?: $_W['setting']['remote']['ftp']['pasv'],
                'timeout' => $_W['setting']['remote']['ftp']['timeout'] ?: $_W['setting']['remote']['ftp']['overtime'],
                'rootdir' => $_W['setting']['remote']['ftp']['rootdir'] ?: $_W['setting']['remote']['ftp']['dir'],
            );
            $ftp = new Ftp($ftp_config);
            if (true === $ftp->connect()) {
                $response = $ftp->upload($filePath, $filename);
                if ($auto_delete_local) {
                    file_delete($filePath);
                }
                if (!empty($response)) {
                    return true;
                } else {
                    return error(1, '远程附件上传失败，请检查配置并重新上传');
                }
            } else {
                return error(1, '远程附件上传失败，请检查配置并重新上传');
            }
        } elseif ($_W['setting']['remote']['type'] == ATTACH_OSS) {
            load()->library('oss');
            load()->model('attachment');
            $buckets = attachment_alioss_buctkets($_W['setting']['remote']['alioss']['key'], $_W['setting']['remote']['alioss']['secret']);
            $host_name = $_W['setting']['remote']['alioss']['internal'] ? '-internal.aliyuncs.com' : '.aliyuncs.com';
            $endpoint = 'http://' . $buckets[$_W['setting']['remote']['alioss']['bucket']]['location'] . $host_name;
            try {
                $ossClient = new \OSS\OssClient($_W['setting']['remote']['alioss']['key'], $_W['setting']['remote']['alioss']['secret'], $endpoint);
                $ossClient->uploadFile($_W['setting']['remote']['alioss']['bucket'], $filename, $filePath);
            } catch (\OSS\Core\OssException $e) {
                return error(1, $e->getMessage());
            }
            if ($auto_delete_local) {
                file_delete($filePath);
            }
        } elseif ($_W['setting']['remote']['type'] == ATTACH_QINIU) {
            load()->library('qiniu');
            $auth = new Qiniu\Auth($_W['setting']['remote']['qiniu']['accesskey'], $_W['setting']['remote']['qiniu']['secretkey']);
            $config = new Qiniu\Config();
            $uploadmgr = new Qiniu\Storage\UploadManager($config);
            // 构造上传策略，覆盖已有文件
            $putpolicy = Qiniu\base64_urlSafeEncode(json_encode(array(
                'scope' => $_W['setting']['remote']['qiniu']['bucket'] . ':' . $filename,
            )));
            $uploadtoken = $auth->uploadToken($_W['setting']['remote']['qiniu']['bucket'], $filename, 3600, $putpolicy);
            list($ret, $err) = $uploadmgr->putFile($uploadtoken, $filename, $filePath);
            if ($auto_delete_local) {
                file_delete($filePath);
            }
            if (null !== $err) {
                return error(1, '远程附件上传失败，请检查配置并重新上传');
            } else {
                return true;
            }
        } elseif ($_W['setting']['remote']['type'] == ATTACH_COS) {
            load()->library('cosv5');
            try {
                $bucket = $_W['setting']['remote']['cos']['bucket'] . '-' . $_W['setting']['remote']['cos']['appid'];
                $cosClient = new Qcloud\Cos\Client(
                    array(
                        'region' => $_W['setting']['remote']['cos']['local'],
                        'credentials'=> array(
                            'secretId'  => $_W['setting']['remote']['cos']['secretid'],
                            'secretKey' => $_W['setting']['remote']['cos']['secretkey'])));
                $cosClient->Upload($bucket, $filename, fopen($filePath, 'rb'));
                if ($auto_delete_local) {
                    file_delete($filePath);
                }
            } catch (\Exception $e) {
                return error(1, $e->getMessage());
            }
        }
        return true;
    }

    public static function MobileUrl($do, $query = array(),$addHost=true, $noredirect = true,$moduleName=ModuleName) {
        $query['do'] = $do;
        $query['m'] = strtolower($moduleName);
        return murl('entry', $query, $noredirect,$addHost);
    }
    public static function WebUrl($do,$moduleName=ModuleName){
        $query['do'] = $do;
        $query['module_name'] = strtolower($moduleName);

        return wurl('site/entry', $query);
    }
    public static function ImgUrl2Site($url){
        $fileName = self::RandomName("tmp",".png");
        $row = file_get_contents($url);
        file_put_contents($fileName,$row);
        load()->func('file');
        $tmpPath = self::TmpPath($fileName);
        return self::TmpImgSave($tmpPath);
    }
    public static function TmpImgSave($tmpPath,$justLocal = false){
        global $_W;
        $fileName = self::RandomName("tmp",".png");
        load()->func('file');
        $result = file_upload(["name"=>$fileName,"tmp_name"=>$tmpPath], 'image');
        if(is_error($result)){
            return $result;
        }
        $url = $_W["attachurl_local"].$result["path"];
        if (!$justLocal){
            $remote = file_remote_upload($result["path"]);
            if(!is_error($remote)){
                $url = $_W["attachurl"].$result["path"];
            }
        }
        file_delete($tmpPath);
        return ["url"=>$url, "path"=>$result["path"]];
    }
    public static function RandomName($pre="",$ext=""){
        return $pre.date('Ymd').random(6).$ext;
    }
    public static function TmpPath($fileName){
        global $_W;
        return IA_ROOT.DIRECTORY_SEPARATOR.explode("/",$_W['script_name'])[1].DIRECTORY_SEPARATOR.$fileName;
    }
}