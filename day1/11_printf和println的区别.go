/*
* @Author: tkk1234
* @Date:   2018-05-27 13:45:36
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 16:00:59
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	a := 10
	//一段段处理，自动换行
	fmt.Println("a = ", a)
	//格式化输出，把a的内容放在%d的位置
	fmt.Printf("a = %d\n", a)

	b, c := 20, 30
	fmt.Println("a = ", a, "b= ", b, "c = ", c) //结论：输出效果一样，但下面书写方便。
	fmt.Printf("a = %d, b = %d，c = %d\n", a, b, c)
}
