/*
* @Author: tkk1234
* @Date:   2018-05-27 17:37:33
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 21:01:50
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	//1，iota叫做常量自动生成器，每隔一行自动累加1
	//2，iota是给常量赋值使用，每隔一行自动累加1

	const (
		a = iota
		b = iota
		c = iota
	)
	// 入口函数 main（就像一个商场的入口）,大括号必须和函数“main”这个英文同行！
	// Println会自动换行
	// 调用函数大部分都需要调用包，即上面的import fmt格式化包。
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c) //go语言结尾没有分号
	//3，iota遇到const就会重置为0
	const d = iota
	fmt.Printf("d value is %d\n", d)

	//4，可以只写一个iota，结果和上面都写一样
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

	//5,如果是同一行，结果都一样
	const (
		j       = iota
		k, l, m = iota, iota, iota //k，l，m结果都是1
		o       = iota
	)
	fmt.Printf("j = %d,k = %d,l = %d,m = %d,o = %d\n", j, k, l, m, o)
}
