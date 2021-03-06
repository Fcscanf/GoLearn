package main

import "fmt"

/*进制转换-位运算*/
func main() {

	//位运算
	/*
		2的补码0000 0010
		3的补码0000 0011
		2&3=》0000 0010==》2
		2|3=》0000 0011==》3
		2^3=》0000 0001==》1
	*/
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)

	/*
		-2的源码 1000 0010=》反码1111 1101=》补码1111 1110
		2的补码                                0000 0010
		-2^2=====补码异或运算，异为1，同为0=====>>1111 1100==》补码
		         补码-1计算出反码=============》1111 1011==》反码
		         反码符号位不变其他位取反得到原码》1000 0100==》原码==》-4
	*/
	fmt.Println(-2 ^ 2)

	/*
		右移运算 ：低位溢出，符号位不变，并用符号位补溢出的高位
		左移运算 ：符号位不变，低位补0
		a:=1>>2  //0000 0001==>0000 0000==>0
		b:=1<<2  //0000 0001==>0000 0100==>4
	*/
	a := 1 >> 2
	b := 1 << 2
	fmt.Println("a=", a, "b=", b)

}
