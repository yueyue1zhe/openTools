<?php
/**
 * bdAiTx2Video.php
 * Created by PhpStorm
 * Author: zhouzhe5934@icloud.com
 * Date  : 2021/3/5
 */

namespace yueyue\service;


class bdAiTx2Video {
    protected $key;
    protected $secret;
    public $Man = 106;
    public $Woman = 5;
    public function __construct($key,$secret) {
        $this->key = $key;
        $this->secret = $secret;
    }
    protected function getToken(){
        $url = "https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id={$this->key}&client_secret={$this->secret}";
        $res = ihttp_get($url);
        if ($res["code"] != 200)return error(1,"获取token请求链接错误");
        $content = json_decode($res["content"],true);
        return $content["access_token"];
    }
    public function Make($str,$per=5,$spd=5,$pit=5,$vol=5){
        $str = trim($str);
        if (empty($str))return error(1,"字符串不能为空");
        $token = $this->getToken();
        if (is_error($token))return $token["message"];
        $url = "https://tsn.baidu.com/text2audio";
        $res = ihttp_post($url,[
            "tex" =>$str,
            "spd" =>$spd,
            "pit" => $pit,
            "vol" => $vol,
            "aue" => 3,
            "per" => $per,
            "tok" => $token,
            "ctp" => 1,
            "lan" => "zh",
            "cuid" => ModuleName
        ]);
        if ($res["code"] != 200)return error(1,"请求链接错误");
        if ($res["headers"]["Content-Type"] == "application/json"){
            $content = json_decode($res["content"],true);
            switch ($content["err_no"]){
                case 500:
                    $err = "不支持输入";
                    break;
                case 501:
                    $err = "输入参数不正确";
                    break;
                case 502:
                    $err = "token验证失败";
                    break;
                case 503:
                    $err = "合成后端错误";
                    break;
                default:
                    $err = "错误";
                    break;
            }
            return error(1,$err.":".$content["err_msg"]);
        }
        if ($res["headers"]["Content-Type"] == "audio/mp3"){
            $filename = date("Ym").md5(time().mt_rand(10, 99)).".mp3";
            file_put_contents($filename,$res["content"]);
            load()->func('file');
            $result = file_upload(["name"=>$filename,"tmp_name"=>IA_ROOT.'/web/'.$filename], 'audio');
            if (!is_error($result))file_remote_upload($result["path"], 'audio');
            file_delete(IA_ROOT.'/web/'.$filename);
            global $_W;
            return ["path"=>$result["path"],"url"=>$_W["attachurl"].$result["path"]];
        }else{
            return error(1,"未知异常");
        }
    }
}