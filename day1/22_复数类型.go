/*
* @Author: tkk1234
* @Date:   2018-05-28 22:44:54
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-29 08:25:46
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	compl := 3.5 + 6i                       //复数complex，3.5为实部，6为虚部，i为虚数单位
	fmt.Printf("compl type is %T\n", compl) //go语言结尾没有分号

	//通过内建函数，取实部和虚部
	//real是求复数的实数部分的函数
	//imag是求复数的虚数部分的函数
	fmt.Println("real(compl) = ", real(compl), "imag(compl) = ", imag(compl))
}
