/*
* @Author: tkk1234
* @Date:   2018-05-27 16:12:52
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 16:56:41
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	// 变量：程序运行期间，可以改变的量，变量声明需要关键字：var
	// 常量：程序运行期间，不可以改变的量，常量声明需要关键字：const
	// 调用函数大部分都需要调用包，即上面的import fmt格式化包。
	const (
		b = 10 //常量赋值。注意，这里没有使用 :=自动推导类型
	)
	fmt.Println(b)                        //go语言结尾没有分号
	fmt.Printf("const b's type is %T", b) //看一下这个常量b的类型
}
