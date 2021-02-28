<?php
/**
 * appcors.php
 * Created by PhpStorm
 * Author: zhouzhe5934@icloud.com
 * Date  : 2021/2/10
 * /app/index.php file
 * require '../framework/bootstrap.inc.php';
 * require_once IA_ROOT."/addons/yueyue_uniapp/lib/wx-cors.php";
 * require IA_ROOT . '/app/common/bootstrap.app.inc.php';
 */

global $_GPC;
$func = ["frontend"];
$dose = $_GPC["m"] == "yueyue_uniapp" && in_array($_GPC["do"],$func);
if($_SERVER['REQUEST_METHOD'] == 'OPTIONS' && $dose) {
    header("Access-Control-Allow-Headers: content-type,token");
    header("Access-Control-Allow-Origin: *");
    header('Access-Control-Allow-Methods: OPTIONS');
    header('Content-type: application/json');
    exit();
}