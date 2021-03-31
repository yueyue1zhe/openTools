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
        if (empty($vendorDir) && defined("PluginMap") && in_array($vendor[1],PluginMap)){
            $vendorDir =  str_replace(ModuleName,ModuleName."_plugin_".$vendor[1],__DIR__);
        }
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

    protected function token_award($data, $jwtUuid,$exp=7200){
        return \Jwt::awardToken($data,$this->api_scope,$jwtUuid,$exp);
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
    public static function ArraySort($array,$keys,$sort='asc') {
        $newArr = $valArr = array();
        foreach ($array as $key=>$value) {
            $valArr[$key] = $value[$keys];
        }
        ($sort == 'asc') ?  asort($valArr) : arsort($valArr);
        reset($valArr);
        foreach($valArr as $key=>$value) {
            $newArr[$key] = $array[$key];
        }
        return $newArr;
    }
    public static function ApiEntryPath(){
        global $_GPC;
        return $_GPC[ApiEntryKey];
    }
    public static function ApiPostKeys(){
        global $_GPC;
        return array_keys($_GPC["__input"]);
    }
    public static function IsPhoneNum($num){
        return preg_match("/^1[34578]\d{9}$/",$num);
    }
    public static function ApiValidate($param){
        if(is_error($param))self::ReqFail($param["message"],$param["errno"]);
        return $param;
    }
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
    public static function setMarCors($allowOrigin){
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
    public static function throwToast($message){
        return error(1,$message);
    }
    public static function throwRedicet($title,$path){
        return error(2, ["title"=>$title,"path"=>$path]);
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
    public static function Arr2sort($arr,$key,$sort=SORT_DESC){
        $last_key = array_column($arr,$key);
        array_multisort($last_key ,$sort,$arr);
        return $arr;
    }
    /**
     * 当前服务器操作系统是否win
     * @return bool
     */
    public static function is_win(){
        return strtoupper(substr(PHP_OS,0,3))==='WIN';
    }

    public static function device_is_ios(){
        return self::get_device_type() == "ios";
    }

    /**
     * 当前客户端类型
     * @return string android || ios || other
     */
    public static function get_device_type(){
        $agent = strtolower($_SERVER['HTTP_USER_AGENT']);
        $type = 'other';
        if(strpos($agent, 'iphone') || strpos($agent, 'ipad')){
            $type = 'ios';
        }
        if(strpos($agent, 'android')){
            $type = 'android';
        }
        return $type;
    }
    public static function TimestampComputed($year, $month, $day, $timestamp=0){
        $dos = sprintf("+%s year +%s month +%s day",$year,$month,$day);
        return  $timestamp ?  strtotime($dos,$timestamp) : strtotime($dos);
    }
    public static function Sec2TimeString($time,$showHMS = true){
        $arr = self::Sec2Time($time);
        $str = "";
        $arr["years"] > 0 ? $str .= $arr["years"]."年": null;
        $arr["days"] > 0 ? $str .= $arr["days"]."天": null;
        if($showHMS){
            $arr["hours"] > 0 ? $str .= $arr["hours"]."个小时": null;
            $arr["minutes"] > 0 ? $str .= $arr["minutes"]."分钟": null;
            $arr["seconds"] > 0 ? $str .= $arr["seconds"]."秒": null;
        }
        return $str;
    }
    public static function Sec2Time($time){
        if(is_numeric($time)){
            $value = array(
                "years" => 0, "days" => 0, "hours" => 0,
                "minutes" => 0, "seconds" => 0,
            );
            if($time >= 31536000){
                $value["years"] = floor($time/31536000);
                $time = ($time%31536000);
            }
            if($time >= 86400){
                $value["days"] = floor($time/86400);
                $time = ($time%86400);
            }
            if($time >= 3600){
                $value["hours"] = floor($time/3600);
                $time = ($time%3600);
            }
            if($time >= 60){
                $value["minutes"] = floor($time/60);
                $time = ($time%60);
            }
            $value["seconds"] = floor($time);
            return (array) $value;
        }else{
            return (bool) FALSE;
        }
    }
    public static function s_to_is($s=0){
        $i    =    floor($s/60);
        $s    =    $s%60;
        $i    =    (strlen($i)==1)?'0'.$i:$i;
        $s    =    (strlen($s)==1)?'0'.$s:$s;
        return $i.':'.$s;
    }

    /**
     * 获取指定时间戳农历日期
     * @param int $timestamp
     * @return string
     */
    public static function getLunar($timestamp){
        $solar = new ComputedSolar();
        $solar->solarYear = date("Y",$timestamp);
        $solar->solarMonth = (int)date("m",$timestamp);
        $solar->solarDay = (int)date("d",$timestamp);
        $lunar = ComputedLunarSolarConverter::SolarToLunar($solar);
        $lunar->isleap ? $leap = "(闰)": $leap = "";
        $time = self::intDate2strDate($lunar->lunarMonth,$lunar->lunarDay);
        return sprintf("农历%s%s",$leap,$time);
    }
    /**
     * 月份转换至农历
     * @param $intM
     * @param $intD
     * @return string
     */
    public static function intDate2strDate($intM,$intD){
        $mArr = ["","正月","二月","三月","四月","五月","六月","七月","八月","九月","十月","冬月","腊月"];
        $dArr = ["","初一","初二","初三","初四","初五","初六","初七","初八","初九","初十", "十一","十二","十三","十四","十五","十六","十七","十八","十九","二十", "廿一","廿二","廿三","廿四","廿五","廿六","廿七","廿八","廿九","三十"];
        return $mArr[$intM].$dArr[$intD];
    }
    /**
     * 获取指定时间戳是星期几
     * @param int $timestamp
     * @return string
     */
    public static function GetWeek($timestamp){
        $weekArray=array("日","一","二","三","四","五","六");
        return "星期".$weekArray[date("w",$timestamp)];
    }
    public static function phpLongApiInit(){
        ignore_user_abort(true);
        set_time_limit(0);
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
        if ($withUniacid && !in_array("uniacid",array_keys($con)))$con["uniacid"]=$_W["uniacid"];
        return pdo_count($this->getTableName(),$con,0);
    }

    public function UniacidNewData($con=[]){
        if (empty($con["uniacid"])){
            global $_W;
            $con["uniacid"] = $_W["uniacid"];
        }
        if (empty($con["createtime"]))$con["createtime"]=TIMESTAMP;
        $param = $this->FillTableField($con);
        pdo_insert($this->getTableName(),$param);
        return pdo_insertid();
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
    public static function DoseModuleStatic($path){
        global $_W;
        return "{$_W['siteroot']}addons/{$_W['current_module']['name']}/{$path}";
    }
    /**
     * 微擎临时带参二维码-时限检测
     * @param string $keyword 关键字
     * @param string $scene 参数
     * @return bool
     */
    public static function FetchWxLinshiQrcodeValid($keyword,$scene){
        global $_W;
        $qrcode = pdo_get("qrcode",array(
            "uniacid" => $_W["uniacid"],
            "type" => "scene",
            "qrcid" => $scene,
            "name" => $keyword,
            "keyword" =>$keyword,
            "model" => 1,
            "status" => 1
        ));
        if (($qrcode["expire"]+$qrcode["createtime"]) < TIMESTAMP){
            pdo_delete("qrcode",array("id"=>$qrcode["id"]));
            return false;
        }
        return true;
    }
    /**
     * 生成临时带参二维码并创建关键字
     * @param string $keyword 关键字
     * @param string $module_name 模块名称
     * @return mixed ['url'] 临时二维码链接 ['qrcid'] 临时二维码携带参数
     */
    public static function MakeWxLinshiQrCodeSyncKeyWord($keyword,$module_name=ModuleName){
        $result = self::MakeWxLinShiQrCode();
        if(is_error($result)){
            return $result;
        }
        self::FetchQrcodeKeyWord($keyword,$result['qrcid'],$result["url"],$result["ticket"]);
        self::FetchKeyWord($keyword,$module_name);
        return $result;
    }
    /**
     * 微擎生成临时带参二维码
     * @return mixed ['qrcid'] 临时二维码携带的参数  ['url']临时二维码链接 ["ticket"]临时二维码票据
     */
    public static function MakeWxLinShiQrCode(){
        $scene = pdo_get("qrcode",[],array("MAX(qrcid) as choose"))["choose"]+1;
        $barcode = array(
            'expire_seconds' => 2592000,
            'action_name' => 'QR_SCENE',
            'action_info' => array(
                'scene' => array(
                    'scene_id' => $scene #临时二维码携带参数
                ),
            ),
        );
        $account_api = WeAccount::create();
        $result = $account_api->barCodeCreateDisposable($barcode);
        if (is_error($result)){
            return $result;
        }
        $result["qrcid"] = $scene;
        return $result;
    }
    /**
     * 微擎生成二维码关键字
     * @param string $keyword 关键字
     * @param string $scene 临时二维码 qrcid MakeWxLinShiQrCode函数中可获取
     * @param string $url 临时二维码 url MakeWxLinShiQrCode函数中可获取
     * @param string $ticket 临时二维码 ticket MakeWxLinShiQrCode函数中可获取
     */
    public static function FetchQrcodeKeyWord($keyword,$scene,$url,$ticket){
        global $_W;
        $qrcodeKeyWord = pdo_get("qrcode",array(
            "uniacid"=>$_W["uniacid"],
            "qrcid" => $scene,
            "keyword"=>$keyword,
            "url" => $url,
            "ticket" => $ticket
        ));
        if (empty($qrcodeKeyWord)){
            pdo_insert("qrcode",array(
                "uniacid" => $_W["uniacid"],
                "acid" => $_W["uniacid"],
                "type" => "scene",
                "extra" => 0,
                "qrcid" => $scene,
                "scene_str" => $keyword,
                "name" => $keyword,
                "keyword" => $keyword,
                "model" =>  1,
                "ticket" => $ticket,
                "url" => $url,
                "expire" => 2592000,
                "subnum" => 0,
                "createtime" => TIMESTAMP,
                "status" => 1
            ));
        }
    }
    public static function W7RouteMobileUrl($method,$query="",$moduleName=ModuleName){
        global $_W;
        return $url = "{$_W["siteroot"]}app/index.php?i={$_W['uniacid']}&c=entry&do={$method}&m=".$moduleName.$query;
    }
    public static function FakeCurl($url,$timeout=1){
        load()->func('communication');
        //todo:: 换成 fsockopen ？
        ihttp_request($url,"", $extra = array(), $timeout);
    }
    public static function searchForm($searchKey,$label,$btn="搜索"){
        global $_W,$_GPC;
        return <<<EOF
<div class="flex-def" >
    <form action="" method="post" style="border: 1px solid #ededed;padding: .5rem 2rem;border-radius: 5px">
        <span>{$label}</span>
        <input id="key_{$searchKey}" type="text" value="{$_GPC[$searchKey]}" style="background-color: #ededed;padding: 5px 10px;border: none;margin: 0 1rem">
        <input id="searchBtn_{$searchKey}" type="button" name="search" value="{$btn}" class="btn btn-primary btn-sm">
    </form>
    <script>
        $("#searchBtn_{$searchKey}").click(()=>{
            let value = $("#key_{$searchKey}").val();     
            let url = window.document.location.href;
            let keyword = "&{$searchKey}=";
            if (url.indexOf(keyword) != -1){
                url = url.substr(0,url.indexOf(keyword))+keyword+value;
            }else {
                url += keyword+value;
            }
            window.location.href = url;
        });
    </script>
</div>
EOF;
    }
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

    /**
     * 微擎生成二维码
     * @param string $content 二维码内容
     * @param bool $justLocal
     * @return array ['path'] ['url']
     */
    public static function MakeQrCode($content,$justLocal = false){
        load()->library("qrcode");
        $fileName = W7Util::RandomName("tmp",".png");
        QRcode::png($content,$fileName,QR_ECLEVEL_H,12,1);
        return self::TmpImgSave(self::TmpPath($fileName),$justLocal);
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
        if (is_error($oauth))return $oauth;
        if (empty($oauth["openid"]))return error(8,"not find openid");
        $oauth["follow"] = $oauth_account->fansQueryInfo($oauth["openid"])["subscribe"];
        if ($oauth["scope"] == "userinfo" || $oauth["scope"] == "snsapi_userinfo"){
            $userinfo = $oauth_account->getOauthUserInfo($oauth['access_token'], $oauth['openid']);
            $oauth["nickname"] = stripslashes(stripcslashes($userinfo['nickname']));
            $oauth["avatar"] = $userinfo['headimgurl'];
            $oauth["unionid"] = $userinfo["unionid"];
            $oauth['gender'] = $userinfo['sex'];
            $oauth['residecity'] = $userinfo['city'] . '市';
            $oauth['resideprovince'] = $userinfo['province'] . '省';
            $oauth['nationality'] = $userinfo['country'];
        }
        self::SetSession($oauth);
        return $oauth;
    }
    public static function SetSession($oauth){
        global $_W;
        $_W['openid'] = $oauth["openid"];
        load()->classs('wesession');
        WeSession::start($_W['uniacid'], $_W['openid']);
        $_SESSION['openid'] = $oauth["openid"];
        $_SESSION['session_key'] = $oauth["session_key"];
        $_SESSION['oauth_openid'] = $oauth["openid"];
        if (empty($userinfo)) {
            $userinfo = array(
                'openid' => $oauth["openid"],
            );
        }
        $_SESSION['userinfo'] = base64_encode(iserializer($userinfo));
        session_id($_W['session_id']);
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
    public static function ImgUrl2Site($url,$justLocal=false){
        $fileName = self::RandomName("tmp",".png");
        $row = file_get_contents($url);
        file_put_contents($fileName,$row);
        load()->func('file');
        $tmpPath = self::TmpPath($fileName);
        return self::TmpImgSave($tmpPath,$justLocal);
    }
    public static function Src2Media($src,$type="images"){
        $account_api = WeAccount::create();
        $result = $account_api->uploadMedia($src, $type);
        return $result;
    }
    public static function Media2Src($media_id){
        $account_api = WeAccount::create();
        $result = $account_api->downloadMedia($media_id);
        return $result;
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

class ComputedLunar{
    public $isleap;
    public $lunarDay;
    public $lunarMonth;
    public $lunarYear;
}

class ComputedSolar{
    public $solarDay;
    public $solarMonth;
    public $solarYear;
}
class ComputedLunarSolarConverter{

    public static $lunar_month_days =
        array(
            1887, 0x1694, 0x16aa, 0x4ad5, 0xab6, 0xc4b7, 0x4ae, 0xa56, 0xb52a,
            0x1d2a, 0xd54, 0x75aa, 0x156a, 0x1096d, 0x95c, 0x14ae, 0xaa4d, 0x1a4c, 0x1b2a, 0x8d55, 0xad4, 0x135a, 0x495d,
            0x95c, 0xd49b, 0x149a, 0x1a4a, 0xbaa5, 0x16a8, 0x1ad4, 0x52da, 0x12b6, 0xe937, 0x92e, 0x1496, 0xb64b, 0xd4a,
            0xda8, 0x95b5, 0x56c, 0x12ae, 0x492f, 0x92e, 0xcc96, 0x1a94, 0x1d4a, 0xada9, 0xb5a, 0x56c, 0x726e, 0x125c,
            0xf92d, 0x192a, 0x1a94, 0xdb4a, 0x16aa, 0xad4, 0x955b, 0x4ba, 0x125a, 0x592b, 0x152a, 0xf695, 0xd94, 0x16aa,
            0xaab5, 0x9b4, 0x14b6, 0x6a57, 0xa56, 0x1152a, 0x1d2a, 0xd54, 0xd5aa, 0x156a, 0x96c, 0x94ae, 0x14ae, 0xa4c,
            0x7d26, 0x1b2a, 0xeb55, 0xad4, 0x12da, 0xa95d, 0x95a, 0x149a, 0x9a4d, 0x1a4a, 0x11aa5, 0x16a8, 0x16d4,
            0xd2da, 0x12b6, 0x936, 0x9497, 0x1496, 0x1564b, 0xd4a, 0xda8, 0xd5b4, 0x156c, 0x12ae, 0xa92f, 0x92e, 0xc96,
            0x6d4a, 0x1d4a, 0x10d65, 0xb58, 0x156c, 0xb26d, 0x125c, 0x192c, 0x9a95, 0x1a94, 0x1b4a, 0x4b55, 0xad4,
            0xf55b, 0x4ba, 0x125a, 0xb92b, 0x152a, 0x1694, 0x96aa, 0x15aa, 0x12ab5, 0x974, 0x14b6, 0xca57, 0xa56, 0x1526,
            0x8e95, 0xd54, 0x15aa, 0x49b5, 0x96c, 0xd4ae, 0x149c, 0x1a4c, 0xbd26, 0x1aa6, 0xb54, 0x6d6a, 0x12da, 0x1695d,
            0x95a, 0x149a, 0xda4b, 0x1a4a, 0x1aa4, 0xbb54, 0x16b4, 0xada, 0x495b, 0x936, 0xf497, 0x1496, 0x154a, 0xb6a5,
            0xda4, 0x15b4, 0x6ab6, 0x126e, 0x1092f, 0x92e, 0xc96, 0xcd4a, 0x1d4a, 0xd64, 0x956c, 0x155c, 0x125c, 0x792e,
            0x192c, 0xfa95, 0x1a94, 0x1b4a, 0xab55, 0xad4, 0x14da, 0x8a5d, 0xa5a, 0x1152b, 0x152a, 0x1694, 0xd6aa,
            0x15aa, 0xab4, 0x94ba, 0x14b6, 0xa56, 0x7527, 0xd26, 0xee53, 0xd54, 0x15aa, 0xa9b5, 0x96c, 0x14ae, 0x8a4e,
            0x1a4c, 0x11d26, 0x1aa4, 0x1b54, 0xcd6a, 0xada, 0x95c, 0x949d, 0x149a, 0x1a2a, 0x5b25, 0x1aa4, 0xfb52,
            0x16b4, 0xaba, 0xa95b, 0x936, 0x1496, 0x9a4b, 0x154a, 0x136a5, 0xda4, 0x15ac
        );

    public static $solar_1_1 =
        array(
            1887, 0xec04c, 0xec23f, 0xec435, 0xec649, 0xec83e, 0xeca51, 0xecc46, 0xece3a,
            0xed04d, 0xed242, 0xed436, 0xed64a, 0xed83f, 0xeda53, 0xedc48, 0xede3d, 0xee050, 0xee244, 0xee439, 0xee64d,
            0xee842, 0xeea36, 0xeec4a, 0xeee3e, 0xef052, 0xef246, 0xef43a, 0xef64e, 0xef843, 0xefa37, 0xefc4b, 0xefe41,
            0xf0054, 0xf0248, 0xf043c, 0xf0650, 0xf0845, 0xf0a38, 0xf0c4d, 0xf0e42, 0xf1037, 0xf124a, 0xf143e, 0xf1651,
            0xf1846, 0xf1a3a, 0xf1c4e, 0xf1e44, 0xf2038, 0xf224b, 0xf243f, 0xf2653, 0xf2848, 0xf2a3b, 0xf2c4f, 0xf2e45,
            0xf3039, 0xf324d, 0xf3442, 0xf3636, 0xf384a, 0xf3a3d, 0xf3c51, 0xf3e46, 0xf403b, 0xf424e, 0xf4443, 0xf4638,
            0xf484c, 0xf4a3f, 0xf4c52, 0xf4e48, 0xf503c, 0xf524f, 0xf5445, 0xf5639, 0xf584d, 0xf5a42, 0xf5c35, 0xf5e49,
            0xf603e, 0xf6251, 0xf6446, 0xf663b, 0xf684f, 0xf6a43, 0xf6c37, 0xf6e4b, 0xf703f, 0xf7252, 0xf7447, 0xf763c,
            0xf7850, 0xf7a45, 0xf7c39, 0xf7e4d, 0xf8042, 0xf8254, 0xf8449, 0xf863d, 0xf8851, 0xf8a46, 0xf8c3b, 0xf8e4f,
            0xf9044, 0xf9237, 0xf944a, 0xf963f, 0xf9853, 0xf9a47, 0xf9c3c, 0xf9e50, 0xfa045, 0xfa238, 0xfa44c, 0xfa641,
            0xfa836, 0xfaa49, 0xfac3d, 0xfae52, 0xfb047, 0xfb23a, 0xfb44e, 0xfb643, 0xfb837, 0xfba4a, 0xfbc3f, 0xfbe53,
            0xfc048, 0xfc23c, 0xfc450, 0xfc645, 0xfc839, 0xfca4c, 0xfcc41, 0xfce36, 0xfd04a, 0xfd23d, 0xfd451, 0xfd646,
            0xfd83a, 0xfda4d, 0xfdc43, 0xfde37, 0xfe04b, 0xfe23f, 0xfe453, 0xfe648, 0xfe83c, 0xfea4f, 0xfec44, 0xfee38,
            0xff04c, 0xff241, 0xff436, 0xff64a, 0xff83e, 0xffa51, 0xffc46, 0xffe3a, 0x10004e, 0x100242, 0x100437,
            0x10064b, 0x100841, 0x100a53, 0x100c48, 0x100e3c, 0x10104f, 0x101244, 0x101438, 0x10164c, 0x101842, 0x101a35,
            0x101c49, 0x101e3d, 0x102051, 0x102245, 0x10243a, 0x10264e, 0x102843, 0x102a37, 0x102c4b, 0x102e3f, 0x103053,
            0x103247, 0x10343b, 0x10364f, 0x103845, 0x103a38, 0x103c4c, 0x103e42, 0x104036, 0x104249, 0x10443d, 0x104651,
            0x104846, 0x104a3a, 0x104c4e, 0x104e43, 0x105038, 0x10524a, 0x10543e, 0x105652, 0x105847, 0x105a3b, 0x105c4f,
            0x105e45, 0x106039, 0x10624c, 0x106441, 0x106635, 0x106849, 0x106a3d, 0x106c51, 0x106e47, 0x10703c, 0x10724f,
            0x107444, 0x107638, 0x10784c, 0x107a3f, 0x107c53, 0x107e48
        );

    public static function GetBitInt($data, $length, $shift)
    {
        return ($data & (((1 << $length) - 1) << $shift)) >> $shift;
    }

    //WARNING: Dates before Oct. 1582 are inaccurate
    public static function SolarToInt($y, $m, $d)
    {
        $m = ($m + 9) % 12;
        $y = intval($y) - intval($m / 10);
        return intval(365 * $y + intval($y / 4) - intval($y / 100) + intval($y / 400) + intval(($m * 306 + 5) / 10) + ($d - 1));
    }

    public static function SolarFromInt($g)
    {
        $y = intval((10000 * intval($g) + 14780) / 3652425);
        $ddd = intval($g - (365 * $y + intval($y / 4) - intval($y / 100) + intval($y / 400)));
        if ($ddd < 0) {
            $y--;
            $ddd = intval($g - (365 * $y + intval($y / 4) - intval($y / 100) + intval($y / 400)));
        }
        $mi = intval((100 * $ddd + 52) / 3060);
        $mm = intval(($mi + 2) % 12 + 1);
        $y = (int)$y + intval(($mi + 2) / 12);
        $dd = intval($ddd - intval(($mi * 306 + 5) / 10) + 1);
        $solar = new ComputedSolar();
        $solar->solarYear = (int)$y;
        $solar->solarMonth = (int)$mm;
        $solar->solarDay = (int)$dd;
        return $solar;
    }

    public static function LunarToSolar($lunar)
    {
        $days = self::$lunar_month_days[$lunar->lunarYear - self::$lunar_month_days[0]];
        $leap = self::GetBitInt($days, 4, 13);
        $offset = 0;
        $loopend = $leap;
        if (!$lunar->isleap) {
            if ($lunar->lunarMonth <= $leap || $leap == 0) {
                $loopend = $lunar->lunarMonth - 1;
            } else {
                $loopend = $lunar->lunarMonth;
            }
        }
        for ($i = 0; $i < $loopend; $i++) {
            $offset += self::GetBitInt($days, 1, 12 - $i) == 1 ? 30 : 29;
        }
        $offset += $lunar->lunarDay;

        $solar11 = self::$solar_1_1[$lunar->lunarYear - self::$solar_1_1[0]];

        $y = self::GetBitInt($solar11, 12, 9);
        $m = self::GetBitInt($solar11, 4, 5);
        $d = self::GetBitInt($solar11, 5, 0);

        return self::SolarFromInt(self::SolarToInt($y, $m, $d) + $offset - 1);
    }

    public static function SolarToLunar($solar)
    {
        $lunar = new ComputedLunar();
        $index = $solar->solarYear - self::$solar_1_1[0];
        $data = ($solar->solarYear << 9) | ($solar->solarMonth << 5) | ($solar->solarDay);
        if (self::$solar_1_1[$index] > $data) {
            $index--;
        }
        $solar11 = self::$solar_1_1[$index];
        $y = self::GetBitInt($solar11, 12, 9);
        $m = self::GetBitInt($solar11, 4, 5);
        $d = self::GetBitInt($solar11, 5, 0);
        $offset = self::SolarToInt($solar->solarYear, $solar->solarMonth, $solar->solarDay) - self::SolarToInt($y, $m, $d);

        $days = self::$lunar_month_days[$index];
        $leap = self::GetBitInt($days, 4, 13);

        $lunarY = $index + self::$solar_1_1[0];
        $lunarM = 1;
        $offset += 1;

        for ($i = 0; $i < 13; $i++) {
            $dm = self::GetBitInt($days, 1, 12 - $i) == 1 ? 30 : 29;
            if ($offset > $dm) {
                $lunarM++;
                $offset -= $dm;
            } else {
                break;
            }
        }
        $lunarD = intval($offset);
        $lunar->lunarYear = $lunarY;
        $lunar->lunarMonth = $lunarM;
        $lunar->isleap = false;
        if ($leap != 0 && $lunarM > $leap) {
            $lunar->lunarMonth = $lunarM - 1;
            if ($lunarM == $leap + 1) {
                $lunar->isleap = true;
            }
        }
        $lunar->lunarDay = $lunarD;
        return $lunar;
    }
}