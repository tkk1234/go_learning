/*
* @Author: tkk1234
* @Date:   2018-05-26 22:45:20
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 13:44:44
 */
package main

import (
	"fmt" //导入fmt包后下面必须要使用，否则报错。
)

func main() {
	//自动推导类型，必须初始化，通过初始化的值确定类型（常用！）
	cc := 30
	// %T打印变量所属的类型
	fmt.Printf("cc type is %T\n", cc)

	//变量，程序运行期间，可以改变的量
	//1，声明的格式为（见下面）：var 变量名 类型
	//2，变量声明后必须要用，否则报错。但只是声明但没有初始化的变量，默认值为0
	//3，同一个{}里声明的变量名是唯一的
	var a int = 20 // 变量的初始化：声明变量的同时，赋值给a为20（一步到位）
	var b, c int   // 还有一种方法：先声明，可以声明多个变量，以逗号分开
	b = 33         //然后再给b赋值
	c = 55         //然后再给c赋值
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
}
