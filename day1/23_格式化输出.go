/*
* @Author: tkk1234
* @Date:   2018-05-29 22:57:29
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-29 23:18:20
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14

	//分别打印各数据格式
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d)

	//指定数据格式输出
	//%d 整型格式
	//%s 字符串格式
	//%c 字符格式
	//%f 浮点格式
	fmt.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)

	//用%v进行自动判断格式类型（但并不是很准确）
	fmt.Printf("a = %v,b = %v,c = %v,d = %v", a, b, c, d)
}
