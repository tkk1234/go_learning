/*
* @Author: harley
* @Date:   2018-06-03 15:49:46
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-03 22:49:23
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

//有参无返回值函数的定义，普通参数列表
func ok(a int) { //括号中的为形参
	//a = 111
	fmt.Println("a = ", a)
}

func ok2(a, b int) { //多个形参，数据类型一样，可以这样写
	fmt.Printf("a= %d, b=%d\n", a, b)
}

func ok3(a, b int, c, d string, e byte) { //多个形参，数据类型一样，可以这样写
	fmt.Printf("a=%d, b=%d,c=%s,d=%s,e=%c\n", a, b, c, d, e)
}

func main() {
	ok(8888)                         //括号中的为实参，只能由实参传递给形参，不能颠倒。
	ok2(99, 66)                      //传实参给ok2的形参
	ok3(152, 478, "you", "are", 'a') //传实参给ok3的形参 %s字符串（用双引号） %c=字符（用单引号）
}
