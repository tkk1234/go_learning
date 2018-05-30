/*
* @Author: tkk1234
* @Date:   2018-05-28 21:39:38
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-28 22:02:22
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	//字符串和字符的区别：
	//1，字符是单引号的，字符串是双引号
	//2，字符只有一个字符，字符串有多个
	var (
		str1 string = "haha" //字符串
	)
	//再来个自动推导类型写法
	str2 := "nonono"

	fmt.Println("str1 = ", str1)
	fmt.Println("str2 = ", str2)

	//测字符串长度
	fmt.Println("str1的字符长度 = ", len(str1))
	fmt.Println("str2的字符长度 = ", len(str2))
}
