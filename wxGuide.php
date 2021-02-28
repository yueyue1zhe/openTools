<?php
/**
 * wxGuide.php
 * Created by PhpStorm
 * Author: zhouzhe5934@icloud.com
 * Date  : 2021/2/28
 */

namespace yueyue\service;


class wxGuide {
    protected $token;

    public function __construct() {
        $this->registerToken();
    }

    public function GetGuideAcctList(){
        $url = "https://api.weixin.qq.com/cgi-bin/guide/getguideacctlist?access_token={$this->token}";
        $res = $this->reqGuide($url,["page"=>0,"num"=>20]);

    }

    /**
     * 获取客户顾问信息
     * @param $openid
     * @return mixed
     */
    public function getGuideBuyerRelationByBuyer($openid){
        $url = "https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelationbybuyer?access_token={$this->token}";
        return $this->reqGuide($url,["openid"=>$openid]);
    }

    /**
     * 为顾问分配客户
     * 每个顾问仅可绑定15000个服务号粉丝
     * 设置用户默认顾问
     * @param $guideOpenid
     * @param $openid
     * @return mixed
     */
    public function AddGuideBuyerRelation($guideOpenid,$openid){
        $url = "https://api.weixin.qq.com/cgi-bin/guide/addguidebuyerrelation?access_token={$this->token}";
        return $this->reqGuide($url,["guide_openid"=>$guideOpenid,"openid"=>$openid]);
    }

    /**
     * 获取顾问信息
     * @param $openid
     * @return mixed
     */
    public function GetGuide($openid){
        $url = "https://api.weixin.qq.com/cgi-bin/guide/getguideacct?access_token={$this->token}";
        return $this->reqGuide($url,["guide_openid"=>$openid]);
    }

    public function DelGuide($openid){
        $url = "https://api.weixin.qq.com/cgi-bin/guide/delguideacct?access_token={$this->token}";
        return $this->reqGuide($url,["guide_openid"=>$openid]);
    }

    /**
     * 添加顾问
     * @param $openid
     * @param string $guideNickname
     * @param string $guideHeadImgUrl
     * @return mixed
     */
    public function AddGuide($openid,$guideNickname = "", $guideHeadImgUrl = "") {
        $url = "https://api.weixin.qq.com/cgi-bin/guide/addguideacct?access_token={$this->token}";
        return $this->reqGuide($url,$this->makeGuidePost($openid,$guideNickname,$guideHeadImgUrl));
    }

    /**
     * 编辑顾问信息
     * @param $openid
     * @param string $guideNickname
     * @param string $guideHeadImgUrl
     * @return mixed
     */
    public function EditGuide($openid, $guideNickname = "", $guideHeadImgUrl = "") {
        $url = "POST https://api.weixin.qq.com/cgi-bin/guide/updateguideacct?access_token={$this->token}";
        return $this->reqGuide($url,$this->makeGuidePost($openid,$guideNickname,$guideHeadImgUrl));
    }

    /**
     * 处理顾问post数据
     * @param $openid
     * @param string $guideNickname
     * @param string $guideHeadImgUrl
     * @return array
     */
    protected function makeGuidePost($openid,$guideNickname = "", $guideHeadImgUrl = ""){
        $con = ["guide_openid" => $openid];
        $guideNickname ? $con["guide_nickname"] = $guideNickname : null;
        //TODO:: 顾问头像，头像url只能用《上传图文消息内的图片获取URL》
        //https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
        $guideHeadImgUrl ? $con["guide_headimgurl"] = $guideHeadImgUrl : null;
        return $con;
    }

    /**
     * post请求顾问接口
     * @param $url
     * @param array $postData
     * @return mixed
     */
    protected function reqGuide($url,$postData=[]){
        $result = ihttp_post($url,json_encode($postData));
        if ($result["code"] != 200){
            return error(1,"网络异常");
        }
        $res = json_decode($result["content"],true);
        if ($res["errcode"] != 0){
            return error($res["errcode"],$this->guideErr($res["errcode"]).':'.$res["errmsg"]);
        }
        return $res;
    }

    /**
     * 顾问接口异常描述
     * @param $errCode
     * @return string
     */
    protected function guideErr($errCode) {
        switch ($errCode) {
            case -1:
                $err = "系统失败";
                break;
            case 40003:
                $err = "无效的openid";
                break;
            case 40032:
                $err = "客户列表大小不合法";
                break;
            case 43004:
                $err = "顾问没有关注该服务号";
                break;
            case 65407:
                $err = "该微信已经绑定为客服，不能再继续绑定顾问";
                break;
            case 9300801:
                $err = "无效的微信号";
                break;
            case 9300802:
                $err = "服务号未开通顾问功能";
                break;
            case 9300803:
                $err = "该微信号已经绑定为顾问";
                break;
            case 9300804:
                $err = "顾问不存在";
                break;
            case 9300806:
                $err = "客户和顾问不存在绑定关系";
                break;
            case 9300810:
                $err = "顾问昵称大于16个字";
                break;
            case 9300812:
                $err = "顾问人数到达上限";
                break;
            default:
                $err = "异常错误,请检查公众号是否已开启【对话能力】";
                break;
        }
        return $err;
    }

    /**
     * 装填token
     */
    protected function registerToken() {
        $account_api = \WeAccount::create();
        $this->token = $account_api->getAccessToken();
    }
}