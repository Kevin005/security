package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

var str1 = "ZHONGGUOnihao123"

func main() {
	//在加密时需要把字符串转换成byte数组类型，解密出来的也是[]bytes类型
	strbytes := []byte(str1)
	//StdEncoding 这个代表是的标准加密，如果使用的是URLEncoding，则是URL加密
	encoded := base64.StdEncoding.EncodeToString(strbytes)
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	//需要把[]bytes类型转换成字符串，使用string进行转换
	decodestr := string(decoded)
	fmt.Println(decodestr)

	// 如果要用在url中，需要使用URLEncoding
	//urleocode只是为了url中一些非ascii字符，可以正确无误的被传输，至于使用哪种编码，就不是eocode所关心和解决的问题了。
	uEnc := base64.URLEncoding.EncodeToString([]byte(str1))
	fmt.Println("URL-----> " + uEnc)
	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("URL-----> " + string(uDec))
}
