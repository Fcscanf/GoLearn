package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*系统常用函数*/

func main() {
	/*汉字先转切片在输出*/
	str := "hello北京"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	/*字符串转整数*/
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}

	/*整数转字符串*/
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str, str)

	/*字符串转[]byte*/
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	/*[]byte转字符串*/
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	/*十进制转二、八、十六进制*/
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是：%v", str)
	strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是：%v", str)

	/*查找子串是否存在指定的字符串*/
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b)

	/*统计一个字符串有几个指定的子串*/
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	/*不区分大小写比较字符串*/
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)

	fmt.Println("结果", "abc" == "Abc")

	/*返回子串在字符串第一次出现的索引值*/
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index=%v\n", index)

	/*返回子串在字符串最后一次出现的索引值*/
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v\n", index)

	/*将指定的字符串替换*/
	str2 := "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v\n,str2=%v\n", str, str2)

	/*按照指定的字符将字符串分隔*/
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	/*将字符串的字母进行大小写转换*/
	str = "golang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)

	/*将字符串左右两边的空格去掉*/
	str = strings.TrimSpace(" tn a lone gopher ntron")
	fmt.Printf("str=%q\n", str)

	/*将字符串左右两边指定的字符去掉*/
	str = strings.Trim("! he!llo!", "!")
	fmt.Printf("str=%q\n", str)

	/*判断字符串是否是以指定的字符串开头*/
	b = strings.HasPrefix("ftp://198.167.324.12", "hsp")
	fmt.Printf("b = %v\n", b)
}
