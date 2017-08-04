function myhash(t,h,y){
	t=t.toString();
	var b = s2b(t);
	if(b==''){
	    b=[];
	}
    if(t.length<17){
        he=[84,132,208,48,203,33,214,6,173,140,172,227,23,205,112,177,173];
        b.push.apply(b,he);
    }
	he = [146,124,150,213,72,186,154,4,126];
	if(h==undefined || h==''){
		h=[];
	}else{
		h=s2b(h.toString());
	}
	
    if(h.length<9){//扰码长度少于8，则自动补充到8位以上以保障复杂程度
		h.push.apply(h,he);
	}
	while(h.length>8){//扰码进行hash至8位
	    h=_h(h);
	}
	while(b.length>16){
		b=_hash(b,h);
		
	}
	var re=[];
	if(y=='32'){
		//32位长度十六进制值输出
		for(var i=0;i<b.length;i++){
			re.push(b[i].toString(16));
		}
		return re.join('').toLowerCase();
	}else{
		for(var i=0;i<b.length;i++){
			re.push(to62(b[i]%62));//简单62进制转换输出
		}
	}
	return re.join('');
}
function to62(b){
	if(b<10){
		return String.fromCharCode(b+48);
	}
	if(b<36){
		return String.fromCharCode(b+55);
	}
	return String.fromCharCode(b+61);
}
//长字串以40个分割为最快计算速度
function _hash(b,h){
	var s=40;//切割分组长度
	var re=[];
	for(var i=0;i<parseInt(b.length/s)+1;i++){
		tmp = b.slice(i*40,(i+1)*40).concat(h);
		re.push.apply(re,dohash(tmp));
	}
	
	return re;
}
//返回16位长度字串符数字数组
function dohash(b){
	while(true){
		b = _h(b);
		if(b.length<17)break;
	}
	return b;
}
/*计算前一个元素与后一个元素相加的结果
 */
function _h(h){
    var hex=h[h[0]%h.length];
	var re=[];
	for(var i=0;i<h.length-1;i++){
		re[i] =(h[i]+h[i+1]+hex)%256;
	}
	return re;
}
/*将字串符转换为16进制数组，即0-255数字数组
 *utf-8会把中文转换为三个元素(utf8中文占用3字节)
 */
function s2b(t) {
		if (!t) {
			return '';
		}
		var utf8 = [];
		for (var i = 0; i < t.length; i++) {
			var s_str = t.charAt(i);
			if (!(/^%u/i.test(escape(s_str)))) {
				//utf8直接解码
				utf8.push(s_str.charCodeAt());
				continue;
			}
			//gbk转utf8
			var s_char = t.charCodeAt(i);
			var b_char = s_char.toString(2).split('');
			var c_char = (b_char.length == 15) ? [0].concat(b_char) : b_char;
			var a_b = [];
			a_b[0] = '1110' + c_char.splice(0, 4).join('');
			a_b[1] = '10' + c_char.splice(0, 6).join('');
			a_b[2] = '10' + c_char.splice(0, 6).join('');
			for (var n = 0; n < a_b.length; n++) {
				utf8.push(parseInt(a_b[n], 2));
			}
		}
		return utf8;
	};
