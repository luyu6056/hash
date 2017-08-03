<?php
function myhash($text,$hmac='',$type='16'){
	$b = s2b($text);
    if(strlen($text)<17){
        $hex=array(84,132,208,48,203,33,214,6,173,140,172,227,23,205,112,177,173);
        $b=array_merge($b,$hex);
    }
	$hex = array(146,124,150,213,72,186,154,4,126);
	$hmac=s2b($hmac);
    if(count($hmac)<9){//扰码长度少于8，则自动补充到8位以上以保障复杂程度
		$hmac=array_merge($hmac,$hex);
	}
	while(count($hmac)>8){//扰码进行hash至8位
	    _h($hmac);
	}
	while(count($b)>16){
		$b=_hash($b,$hmac);
		
	}
	$result='';
	if($type=='32'){
		//32位长度十六进制值输出
		for($i=0;$i<count($b);$i++){
			$result.=chr($b[$i]);
		}
		$result = bin2hex($result);
	}else{
		for($i=0;$i<count($b);$i++){
			$result.=gmp_strval($b[$i]%62,62);//简单62进制转换输出
		}
	}
	return $result;
}
//长字串以40个分割为最快计算速度
function _hash($b,$h){
    
	$s=40;//切割分组长度
	$result=array();
	for($i=0;$i<floor(count($b)/$s)+1;$i++){
		$tmp=array_merge(array_slice($b,$i*$s,$s),$h);
		$result=array_merge($result,dohash($tmp));
	}
	
	return $result;
}
//返回16位长度字串符数字数组
function dohash($b){
	while(true){
		_h($b);
		if(count($b)<17)break;
	}
	return $b;
}
/*指针计算前一个元素与后一个元素相加的结果
 *传入字串符数字数组，会删除最后一个元素
 */
function _h(&$h){
    $hex=$h[$h[0]%count($h)];
	for($i=0;$i<count($h)-1;$i++){
		$h[$i] =($h[$i]+$h[$i+1]+$hex)%256;
	}
	unset($h[$i]);
}
/*将字串符转换为16进制数组，即0-255数字数组
 *utf-8会把中文转换为三个元素(utf8中文占用3字节)
 */
function s2b($text){
    $result=array();
    for($i=0;$i<strlen($text);$i++){
       $result[]=ord(substr($text,$i,1));
    }
    return $result;
}
?>
