/*
* @Author: tkk1234
* @Date:   2018-05-26 22:45:20
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-26 23:00:32
 */
package main

import (
	"fmt" //导入fmt包后下面必须要使用，否则报错。
)

func main() {
	//变量，程序运行期间，可以改变的量
	//1，声明的格式为（见下面）：var 变量名 类型
	//2，变量声明后必须要用，否则报错。但只是声明但没有初始化的变量，默认值为0
	//3，同一个{}里声明的变量名是唯一的
	var a int    //
	var b, c int // 可以声明多个变量，以逗号分开
	a = 10       //给a变量赋值
}
