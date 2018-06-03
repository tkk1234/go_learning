/*
* @Author: tkk1234
* @Date:   2018-05-29 23:19:29
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-29 23:47:08
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	var a int //声明变量
	fmt.Println("请输入变量a：")
	fmt.Scan(&a) //fmt.Scan 阻塞，等待用户输入
	fmt.Println("a = ", a)
}
