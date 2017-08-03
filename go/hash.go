package libraries

import (
	"reflect"
	"fmt"
	"strings"
	"math"
	"unsafe"
)
type Myhash struct {
	Hmac string
}

func Newhash()*Myhash{
	return new(Myhash)
}

const CODE62  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const CODE_LENTH = 62
var EDOC = map[string]int{"0":0,"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"A":10,"B":11,"C":12,"D":13,"E":14,"F":15,"G":16,"H":17,"I":18,"J":19,"K":20,"L":21,"M":22,"N":23,"O":24,"P":25,"Q":26,"R":27,"S":28,"T":29,"U":30,"V":31,"W":32,"X":33,"Y":34,"Z":35,"a":36,"b":37,"c":38,"d":39,"e":40,"f":41,"g":42,"h":43,"i":44,"j":45,"k":46,"l":47,"m":48,"n":49,"o":50,"p":51,"q":52,"r":53,"s":54,"t":55,"u":56,"v":57,"w":58,"x":59,"y":60,"z":61, }
/**
 * 编码 整数 为 base62 字符串
 */
func Base62_Encode(number int) string {
    if number == 0 {
        return "0"
    }
    result := make([]byte , 0)
    for number > 0 {
        round  := number / CODE_LENTH
        remain := number % CODE_LENTH
        result = append(result,CODE62[remain])
        number  = round
    }
    return string(result)
}


/**
 * 解码字符串为整数
 */
func Base62_Decode(str string) int {
    str = strings.TrimSpace(str)
    var result int = 0
    for index,char := range []byte(str){
        result +=  EDOC[string(char)] * int(math.Pow(CODE_LENTH,float64(index)))
    }
    return result
}
/* 参1 需要hash的原始string
 * 参2 混淆值
 * 参3 输出方式,16位六十二进制或者32位十六进制
 */
func (h *Myhash)Hash(text...string)string{
	bb := h.S2B(&text[0])
	b:=make([]byte,len(bb))
	copy(b,bb)
	if(len(b)<17){
		hex := []byte{84,132,208,48,203,33,214,6,173,140,172,227,23,205,112,177,173}
		b=append(b,hex...)
	}
	hex := []byte{146,124,150,213,72,186,154,4,126}
	var hmac string
	if(len(text)==1){
		hmac = "34819d7be"
	}else{
		hmac = text[1]
	}
	out_type := "16"
	if(len(text)==3 && text[2]=="32"){
		out_type = "32"
	}
	hma := h.S2B(&hmac)
	hm:=make([]byte,len(hma))
	copy(hm,hma)
	if(len(hm)<9){
		hm=append(hm,hex...)
	}
	for true{
		_h(&hm)
		if(len(hm)<9){break}
	}
	for true{
		b=_hash(b,hm)
		if(len(b)<17){break}
	}
	
	if(out_type == "32"){
		result:=make([]string,16)
		for key,value := range b{
			result[key]=fmt.Sprintf("%02x", value)
		}
		return strings.Join(result,"")
	}else{
		result:=make([]byte,16)
		for key,value := range b{
			result[key]=CODE62[value % CODE_LENTH]
		}
		return string(result)
	}
	return ""
}
func _hash(b []byte,h []byte)[]byte{
	c:=make([]byte,len(b))
	copy(c,b)
	tmp := []byte{}
	result:=[]byte{}
	s:=40

	for i:=0;i<len(c)/s+1;i++{
		copy(c,b)
		if(i*s+s>len(c)){
			tmp=c[i*s:len(c)]
		}else{
			tmp=c[i*s:i*s+s]
		}
		tmp=append(tmp,h...)
		result=append(result,__hash(tmp)...)
	}
	
	return result
}
func __hash(b []byte)[]byte{
	for true{
		_h(&b)
		if(len(b)<17){break}
	}
	
	return b
}

func _h(b *[]byte){
	hex :=(*b)[int((*b)[0])%len(*b)]
	tmp :=make([]byte,len(*b)-1)
	for i:=0;i<len(*b)-1;i++{
		tmp[i] =(*b)[i]+(*b)[i+1]+hex
	}
	*b=tmp
}
func (h *Myhash)S2B(s *string) []byte {
    return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}

func (h *Myhash)B2S(buf []byte) string {
    return *(*string)(unsafe.Pointer(&buf))
}


func init(){
	//fmt.Println(Newhash().Hash("123","123","32"))
  //81a5f39275fd8f97a79494df034ff6db
  //fmt.Println(Newhash().Hash("123","123","16"))
  //5fvMt5JRhOOb3HyX
}
