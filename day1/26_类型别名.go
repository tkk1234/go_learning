/*
* @Author: tkk1234
* @Date:   2018-05-30 22:54:05
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-30 23:32:33
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	//给int64起一个别名叫hahanono
	type hahanono int64
	var (
		a hahanono
	)
	fmt.Printf("a type is %T\n", a)

	type (
		xxx int64
		yyy byte //注意byte是单个字符，如果下面的a变为多个字符会报错！
	)

	var (
		b xxx = 15 //b为变量名 xxx为类型别名
		c yyy = 'a'
	)

	fmt.Printf("b type is %d,c type is %c\n", b, c)
}
