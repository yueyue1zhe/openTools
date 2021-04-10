<?php
/**
 * api.php
 * Created by PhpStorm
 * Author: zhouzhe5934@icloud.com
 * Date  : 2021/4/10
 */

define('IN_MOBILE', true);
require '../../../framework/bootstrap.inc.php';
global $_W,$_GPC;
$_GPC["do"] = "frontend";
$_GPC["m"] = "yueyue_sellv2";
$urls = parse_url($_W['siteroot']);
$_W['siteroot'] = $urls['scheme'] . '://' . $urls['host']."/";
if ($_SERVER['REQUEST_METHOD'] == 'OPTIONS'){
    header("Access-Control-Allow-Headers: content-type,token");
    header("Access-Control-Allow-Origin: *");
    header('Access-Control-Allow-Methods: OPTIONS');
    header('Content-type: application/json');
    exit();
}
require IA_ROOT . '/app/common/bootstrap.app.inc.php';
if ($_W['settingTable']['copyright']['status'] == 1) {
    $_W['siteclose'] = true;
    message('抱歉，站点已关闭，关闭原因：' . $_W['settingTable']['copyright']['reason']);
}
require IA_ROOT . "/app/source/entry/__init.php";
require IA_ROOT . '/app/source/entry/site.ctrl.php';