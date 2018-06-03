/*
* @Author: tkk1234
* @Date:   2018-06-03 15:35:42
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-03 15:49:07
 */
package main

import (
	"fmt"
)

//无参无返回值函数的定义

func haha() {
	a := 10
	fmt.Println("this is value is:", a)
}

func main() {
	haha() //无参无返回值函数的调用(main函数是入口一定要有)，直接写函数名+（）
}
