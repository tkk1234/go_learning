/*
* @Author: tkk1234
* @Date:   2018-05-27 21:01:41
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 22:01:33
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	// 入口函数 main（就像一个商场的入口）,大括号必须和函数“main”这个英文同行！
	// Println会自动换行
	// 调用函数大部分都需要调用包，即上面的import fmt格式化包。
	var (
		test float32 = 3.1415926
	)
	fmt.Println("test value is", test) //go语言结尾没有分号

	//再来一个自动推导数据类型
	test1 := 3.25614
	fmt.Println("test1 value is", test1)
	fmt.Printf("and test1 data type is %T\n", test1)
	//float64存储小数比float32更准确
}
