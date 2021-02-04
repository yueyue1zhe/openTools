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

class AppUtil{
    private static function getMarRF($scope){
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
        if (!empty($allowOrigin)){
            $origin = isset($_SERVER['HTTP_ORIGIN'])? $_SERVER['HTTP_ORIGIN'] : '';
            if (in_array($origin,$allowOrigin)){
                header("access-control-allow-headers: token,content-type");
                header("access-control-allow-methods: post");
                header("access-control-allow-origin: *");
            }
        }
    }
    public static function Mar($ApiScope,$allowOrigin=[]){
        self::setMarCors($allowOrigin);
        $cf = self::getMarRF($ApiScope);
        if(!class_exists($cf["route"])){
            self::ReqFail("内部错误:路由不存在");
        }
        $apiClass = new $cf["route"]();
        if(!method_exists($apiClass, $cf["func"])){
            self::ReqFail("内部错误:方法不存在");
        }
        $result = call_user_func_array([$apiClass, $cf["func"]], []);
        if(is_error($result)){
            self::ReqFail($result["message"],$result["errno"]);
        }
        self::ReqOk($result);
    }

    public static function MakeQrCode2Show($content){
        require_once IA_ROOT . "/addons/".ModuleName."/lib/exter/phpqrcode.php";
        QRcode::png($content,false,QR_ECLEVEL_H,12,1);
    }
    public static function ReqOk($data){
        self::Req($data,"",0);
    }
    public static function ReqFail($message,$errno=1){
        self::Req("",$message,$errno);
    }
    public static function Req($data,$message,$errno){
        echo json_encode([
            "data" => $data,
            "message" => $message,
            "errno" => $errno
        ]);
        exit();
    }
    public static function UrlIsImage($url){
        return !empty($url) && stripos(get_headers($url)[1],"image") !== false;
    }
    public static function FastModulePathBase($path){
        return IA_ROOT.DIRECTORY_SEPARATOR."addons".DIRECTORY_SEPARATOR.ModuleName.$path;
    }
}

class W7Util{
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