/*
* @Author: tkk1234
* @Date:   2018-05-31 22:05:00
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-01 22:53:31
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	if a := 10; a == 12 {
		fmt.Println("yes")
	} else { //else 格式一定要这样写仅跟在后面
		fmt.Println("no")
	}

	//else if的用法（这种用法执行效率高，因为只要有条件满足就停止运行了）
	if c := 100; c == 200 {
		fmt.Println("相等")
	} else if c > 200 {
		fmt.Println("大于")
	} else if c < 200 {
		fmt.Println("c小于200")
	} else {
		fmt.Println("没办法了，比较不出来了......")
	}

	//都写if的用法执行效率低，因为不管匹不匹配条件都要走一遍
	b := 100 //这里初始化变量声明要单独写出来，否则报错，因为下面的b > 200也需要b := 100的声明
	if b == 200 {
		fmt.Println("相等")
	}
	if b > 200 {
		fmt.Println("大于")
	}

	if b < 200 {
		fmt.Println("b小于200")
	} else {
		fmt.Println("没办法了，比较不出来了......")
	}
}
