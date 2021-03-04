<?php
/**
 * iCash.php
 * Created by PhpStorm
 * Author: zhouzhe5934@icloud.com
 * Date  : 2021/3/4
 */

namespace yueyue\service;


class iCash{
    protected $account;
    protected $payment;
    protected $cert;
    protected $key;
    protected $w;
    public function __construct() {
        global $_W;
        $this->w = $_W;
        $this->account = pdo_fetch("SELECT * FROM " . tablename('account_wechats') . " where uniacid=" . $this->w['uniacid'] . " limit 1");
        $uni_set = pdo_fetch("SELECT * FROM " . tablename('uni_settings') . " where uniacid=" . $this->w['uniacid'] . " limit 1");
        $this->payment = unserialize($uni_set['payment']);
        #绝对路径
        $this->cert = IA_ROOT.'/addons/'.ModuleName.'/cert/'.md5(base64_encode($this->w['uniacid'].'apiclient_cert.pem')).'.pem';
        $this->key = IA_ROOT.'/addons/'.ModuleName.'/cert/'.md5(base64_encode($this->w['uniacid'].'apiclient_key.pem')).'.pem';
    }
    /*
    $amount 发送的金额（分）目前发送金额不能少于1元
    $re_openid, 发送人的 openid
    $desc  //  企业付款描述信息 (必填)
    $check_name    收款用户姓名 (选填)
    */
    public function sendMoney($amount, $re_openid, $desc = '测试', $check_name = '') {
        $total_amount = (100) * $amount;
        $data = array('mch_appid' => $this->account['key'],//商户账号appid
            'mchid' => $this->payment['wechat']['mchid'],//商户号
            'nonce_str' => $this->createNoncestr(),//随机字符串
            'partner_trade_no' => date('YmdHis') . rand(1000, 9999),//商户订单号
            'openid' => $re_openid,//用户openid
            'check_name' => 'NO_CHECK',//校验用户姓名选项,
            're_user_name' => $check_name,//收款用户姓名
            'amount' => $total_amount,//金额
            'desc' => $desc,//企业付款描述信息
            'spbill_create_ip' => '0.0.0.0',//Ip地址
        );
        $secrect_key = $this->payment['wechat']['signkey'];///这个就是个API密码。MD5 32位。
        $data = array_filter($data);
        ksort($data);
        $str = '';
        foreach ($data as $k => $v) {
            $str .= $k . '=' . $v . '&';
        }
        $str .= 'key=' . $secrect_key;
        $data['sign'] = md5($str);
        $xml = $this->arraytoxml($data);

        $url = 'https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers'; //调用接口
        $res = $this->curl($xml, $url);
        $return = $this->xmltoarray($res);
        //返回来的结果
        $responseObj = simplexml_load_string($res, 'SimpleXMLElement', LIBXML_NOCDATA);
        $res = $responseObj->return_code;  //SUCCESS  如果返回来SUCCESS,则发生成功，处理自己的逻辑
        if ($return["return_code"] == "SUCCESS"){
            //连接成功
            if ($return["result_code"] ==  "FAIL"){
                return error(1,"{$return["return_msg"]}:{$return["err_code_des"]}");
            }
            return $return;
        }else{
            return error(1,"请求失败,请检查证书配置");
        }
    }

    protected function createNoncestr($length = 32) {
        $chars = "abcdefghijklmnopqrstuvwxyz0123456789";
        $str = "";
        for ($i = 0; $i < $length; $i++) {
            $str .= substr($chars, mt_rand(0, strlen($chars) - 1), 1);
        }
        return $str;
    }

    protected function unicode() {
        $str = uniqid(mt_rand(), 1);
        $str = sha1($str);
        return md5($str);
    }
    protected function arraytoxml($data) {
        $str = '<xml>';
        foreach ($data as $k => $v) {
            $str .= '<' . $k . '>' . $v . '</' . $k . '>';
        }
        $str .= '</xml>';
        return $str;
    }

    protected function xmltoarray($xml) {
        //禁止引用外部xml实体
        libxml_disable_entity_loader(true);
        $xmlstring = simplexml_load_string($xml, 'SimpleXMLElement', LIBXML_NOCDATA);
        $val = json_decode(json_encode($xmlstring), true);
        return $val;
    }

    protected function curl($param = "", $url) {
        $ch = curl_init();                                      //初始化curl
        curl_setopt($ch, CURLOPT_URL, $url);                 //抓取指定网页
        curl_setopt($ch, CURLOPT_HEADER, 0);                    //设置header
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);            //要求结果为字符串且输出到屏幕上
        curl_setopt($ch, CURLOPT_POST, 1);                      //post提交方式
        curl_setopt($ch, CURLOPT_POSTFIELDS, $param);           // 增加 HTTP Header（头）里的字段
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);        // 终止从服务端进行验证
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);
        curl_setopt($ch, CURLOPT_SSLCERT, $this->cert); //这个是证书的位置绝对路径
        curl_setopt($ch, CURLOPT_SSLKEY, $this->key); //这个也是证书的位置绝对路径
        $data = curl_exec($ch);                                 //运行curl
        curl_close($ch);
        return $data;
    }
}