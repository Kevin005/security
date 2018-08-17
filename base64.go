package main

import (
	"encoding/base64"
	"fmt"
)

var str1 = "ZHONGGUOnihao123"
var str = "abcdef"

func main() {
	//在加密时需要把字符串转换成byte数组类型，解密出来的也是[]bytes类型
	strbytes := []byte(str)
	//StdEncoding 这个代表是的标准加密，如果使用的是URLEncoding，则是URL加密
	encoded := base64.StdEncoding.EncodeToString(strbytes)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	//需要把[]bytes类型转换成字符串，使用string进行转换
	decodestr := string(decoded)
	fmt.Println(decodestr)
}
