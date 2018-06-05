/*
* @Author: tkk1234
* @Date:   2018-06-05 08:04:21
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-05 12:32:17
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

//无参有返回值（return 666）
//给返回值起个变量名（这里写了res），go推荐写法
func test1() (res int) { //返回什么类型的数据就要写相应的数据类型
	res = 66666
	return
}

func main() {
	a := test1() //赋值给a，看函数test1是不是返回了66666
	fmt.Println("a=", a)
}
