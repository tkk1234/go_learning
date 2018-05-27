/*
* @Author: iphone5s
* @Date:   2018-05-27 16:58:54
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 17:34:43
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	// 不同类型变量的声明（定义）
	var (
		a = 55   //自动推导数据类型
		b = 67.8 //自动推导数据类型
	)
	//当然如下这种也可以，但我觉得麻烦
	// a int     = 55   声明a的类型为整数，然后赋值（变量初始化一步到位）
	// b float64 = 67.8 声明b的类型为小数，然后赋值（变量初始化一步到位）
	//
	//a, b = 55, 67.8   //这样分开赋值也可以，看个人。
	fmt.Println(a, b) //go语言结尾没有分号

	//不同类型常量的声明（定义）,可以自动推导数据类型，所以不用声明类型。
	const (
		c = 66
		d = 33.6
	)
	fmt.Println(c, d)
}
