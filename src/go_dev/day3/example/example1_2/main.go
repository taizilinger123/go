package main

import(
	"fmt"
	"strings"
)

func urlProcess(url string) string {
	 
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func main(){
	var(
		url string 
		path string 
	)

	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Println(url)
	fmt.Println(path)
}

/*
G:\project>go run "g:\PROJECT\src\go_dev\day3\example\example1_2\main.go"
10.100.8.8/user/login  c:/study
http://m0.100.8.8/user/login
c:/study/
==============================================
strings.HasPrefix  判断字符串是否以prefix开头
strings.HasSuffix   判断字符串是否以suffix结尾
strings.Index       判断字符串首次出现的位置
strings.LastIndex  判断字符串最后出现的位置
strings.Replace    字符串替换
strings.Count      字符串计数
strings.Repeat     重复count次str
strings.ToLower    转为小写
strings.ToUpper    转为大写
strings.TrimSpace  去掉字符串首尾空白字符
strings.Trim       去掉字符串首尾指定字符
strings.TrimLeft   去掉字符串首指定的字符
strings.TrimRight  去掉字符串尾指定的字符
strings.Field      返回字符串空格的所有子串slice 
strings.Split      返回字符串split分割的所有子串的slice 
strings.Join       连接指定的字符的所有字符串
strings.Itoa       把一个整数转成字符串
strings.Atoi       把一个字符串转成整数
*/