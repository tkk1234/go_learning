/*
* @Author: tkk1234
* @Date:   2018-06-05 12:35:05
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-05 12:43:57
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

//多个返回值
func you() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

func main() {
	//函数调用
	a, b, c := you()
	fmt.Printf("a = %d,b = %d,c = %d\n", a, b, c)
}
