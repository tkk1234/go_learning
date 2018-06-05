/*
* @Author: tkk1234
* @Date:   2018-06-04 22:47:09
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-04 23:19:13
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func test1(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data=", data)
	}
}

func test3(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data=", data)
	}
}

func test2(args ...int) {
	//把全部元素传递给上面的test1
	test1(args...)
	fmt.Println("===========================")
	//只想把后2个参数传递给另外一个函数使用：
	test3(args[2:]...) //从args[2]开始（包括本身），把后面所有元素传递给上面的函数test3
}

func main() {
	test2(1, 2, 3, 4) //传递实参给上面test2的(args ...int)

}
