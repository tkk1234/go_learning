/*
* @Author: tkk1234
* @Date:   2018-06-01 21:54:15
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-01 22:53:08
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	num := 5
	switch num {
	//switch后面写的变量num，当然也可以写成这种形式 switch num := 5;num
	//不像if后面写的是判断条件，运行原理：num和1开始做比较，匹配到哪个就停止匹配其余条件。
	case 1: //也可以写多个条件 1,3,5,6:
		fmt.Printf("现在电梯在%d楼啦~~~", num)
		//默认下面都有一个break跳出语句，不写代表默认就包含，所以不用写
		//下面如果跟 fallthrough语句，代表无条件执行，只要满足一个条件后面的
		//语句都会无条件执行。而不是像break一样停止执行。
	case 2:
		fmt.Printf("现在电梯在%d楼啦~~~", num)
	case 3:
		fmt.Printf("现在电梯在%d楼啦~~~", num)
	case 4:
		fmt.Printf("现在电梯在%d楼啦~~~", num)
	case 5:
		fmt.Printf("现在电梯在%d楼啦~~~", num)
	default: //如果以上都不匹配，就执行下面的语句
		fmt.Println("现在按的是100楼啦，你疯了吗？！")
	}

	//还有一个种方式：
	score := 900
	switch { //不写条件
	case score == 900: //case后面可以放条件
		fmt.Printf("优秀%d楼啦~~~", score)
	case score < 90:
		fmt.Printf("差%d楼啦~~~", score)
	case score == 30:
		fmt.Printf("极差%d楼啦~~~", score)
	default: //如果以上都不匹配，就执行下面的语句
		fmt.Println("0分！你疯了吗？！")
	}
}
