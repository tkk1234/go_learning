/*
* @Author: tkk1234
* @Date:   2018-05-27 20:37:13
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 20:59:18
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	var (
		a bool = true
	)
	// 入口函数 main（就像一个商场的入口）,大括号必须和函数“main”这个英文同行！
	// Println会自动换行
	// 调用函数大部分都需要调用包，即上面的import fmt格式化包。
	fmt.Println("a = ", a) //go语言结尾没有分号
	//自动推导类型
	var (
		b = false
	)

	fmt.Println("b = ", b)
	//或者另外一种自动推导，效果和上面的自动推导一样
	c := false

	fmt.Println("c = ", c)

}
