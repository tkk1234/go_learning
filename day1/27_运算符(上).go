/*
* @Author: tkk1234
* @Date:   2018-05-31 08:07:04
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-31 12:36:57
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {

	fmt.Println("6 > 1比较结果：", 6 > 1) //结果返回布尔值
	fmt.Println("5是不是不等于3？结果：", 5 != 3)
	fmt.Println("6 > 1比较结果：", !(6 > 1))    //结果取反
	fmt.Println("5是不是不等于3？结果：", !(5 != 3)) //结果取反
	// && 与，并且，左右2边都为真，结果才为真，有一个假就全为假。
	fmt.Println("2个为真才为真：", true && true)
	fmt.Println("只要有1个为false就为false：", false && true)
	//|| 或者，左右2边都为假，结果才为假。有一个真就全为真。
	fmt.Println("2个为假才为假：", false && false)
	fmt.Println("只要有1个为true就为true：", true || false)

	a := 8
	fmt.Println("8是否在0和10之间呢？布尔判断结果为：", a > 0 && a < 100)
	//在go语言里要写成a > 0 && a < 100 不能像其他语言一样写成a > 0 < 100
}
