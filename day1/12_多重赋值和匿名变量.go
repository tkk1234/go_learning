/*
* @Author: tkk1234
* @Date:   2018-05-27 14:11:20
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 14:42:14
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func test() (a, b, c int) {
	return 1, 2, 3 //函数返回值
}

func main() {
	i, j := 10, 20
	i, j = j, i //交换2个变量的值i变j，j变i
	fmt.Println(i, j)

	//匿名变量，丢弃数据不处理。匿名变量需配合函数返回值使用，才有优势。
	var c, d, e int
	c, d, e = test() //用上面的函数test return值1, 2, 3赋值给c d e
	fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)

	_, d, _ = test()          //用上面的函数test return值1, 2, 3赋值给c d e，但这里用了匿名变量！
	fmt.Printf("d = %d\n", d) //经常这样用！
}
