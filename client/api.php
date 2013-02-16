<?
$begin = microtime(TRUE);
//-- 统计IP PV --
/*
$pathArr = array('api', 'ip');
$refer = empty($_SERVER['refer']) ? '' : $_SERVER['refer'];
$paraArr = array('adid' => '1', 'ip' => ip2long('202.96.209.133'), 'time' => date('Y-m-d'), 'type' => 0, 'domain' => 'www.q-cms.cn', 'refer' => $refer, 'ua' => $_SERVER['HTTP_USER_AGENT']);
*/
//-- 统计Reg --
/*
$pathArr = array('api', 'reg');
$paraArr = array('adid' => '1', 'ip' => ip2long('202.96.209.133'), 'uid' => 1234, 'time' => date('Y-m-d'));
*/
//-- 统计订单 --
/*
$pathArr = array('api', 'order');
$paraArr = array('adid' => '1', 'uid' => 1234, 'order_id' => 123456, 'money' => '99.01', 'real_money'=> '80.99', 'time' => time());
*/
//-- 邮件统计 --
/*
$pathArr = array('api', 'mail');
$paraArr = array('uid' => 1234, 'mail_id' => '1', 'goods_id' => 123456, 'ip' => ip2long('202.96.209.133'), 'time' => time());
*/
//-- 活动统计 --
/*
$pathArr = array('api', 'activity');
$paraArr = array('uid' => 1234, 'activity_id' => '1', 'goods_id' => 123456, 'ip' => ip2long('202.96.209.133'), 'time' => time());
*/
$result = api($pathArr, $paraArr);
if(empty($result)){
	die('Api connection fail !');
}
echo $result['httpCode'];

function api($pathArr, $paraArr){
	$token = 'young123';
	$secret = "8aa2a6ae0b2a0e6c9390b8e3f6626b47";
	$url = "http://127.0.0.1:12345/".implode('/', $pathArr);
	$ch = curl_init($url);
	curl_setopt($ch, CURLOPT_HTTPHEADER, array('Token:'.$token, 'Secret:'.$secret));
	curl_setopt($ch, CURLOPT_POST, true);		
	curl_setopt($ch, CURLOPT_POSTFIELDS, $paraArr);
	curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);		
	$result = curl_exec($ch);
	$info = curl_getinfo($ch);
	curl_close($ch);
	return json_decode($result, true);
}

$end = microtime(TRUE);
$time = $end-$begin;
echo '<br>此次共花了'.$time.'秒时间。<br>';