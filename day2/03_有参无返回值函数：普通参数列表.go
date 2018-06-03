/*
* @Author: harley
* @Date:   2018-06-03 15:49:46
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-03 16:58:24
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

//有参无返回值函数的定义，普通参数列表
func ok(a int) { //
	//a = 111
	fmt.Println("a = ", a)
}

func main() {
	ok(8888) //括号中成为实参
}
