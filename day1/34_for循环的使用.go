/*
* @Author: tkk1234
* @Date:   2018-06-03 08:42:44
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-03 11:48:01
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		//1, i := 1 初始化条件 ;
		//2, 1 <= 100 判断条件是否为真，如果为真，执行循环。如果为假，跳出循环；
		//3, 条件变化 i++
		//4, 重复2,3,4
		sum = sum + i
	}
	fmt.Println("sum = ", sum)
}
