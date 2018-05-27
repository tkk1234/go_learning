/*
* @Author: tkk1234
* @Date:   2018-05-27 12:58:11
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 13:44:48
 */
package main

import (
	"fmt"
)

func main() {
	var a int // 先声明变量
	a = 10    // 然后再赋值  或者var a int = 10 声明+赋值一步到位
	fmt.Println(a)

	// := 表示自动推导类型，作用：先声明b的类型，再给b赋值为20.这样就比上面的方法快（常用）。
	b := 20
	fmt.Println("b = ", b)

	// 下面如果再用赋值一个同样的自动推导类型 b := 50 就会报错，
	// 因为前面已经有变量b:= 20的自动推导类型
	b = 55 //如果只是单纯赋值是可以的（虽然都是b）
	fmt.Println("b2 = ", b)
}
