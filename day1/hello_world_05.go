/*
* @Author: tkk1234
* @Date:   2018-05-26 21:58:35
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-26 21:59:05
 */
//1）o语言是以包作为管理单位
//2）每个文件必须先声明包
//3）程序必须有一个main包（重要）
package main // mian包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	// 入口函数 main（就像一个商场的入口）,大括号必须和函数main同行！
	// Println会自动换行
	// 调用函数大部分都需要调用包，即上面的import fmt格式化包。
	fmt.Println("hello world!") //go语言结尾没有分号
	fmt.Println("...")
}

/*

这是块注释 只要这区间里的都是注释

*/
