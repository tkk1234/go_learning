/*
* @Author: tkk1234
* @Date:   2018-05-31 12:38:08
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-31 22:05:15
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	s := "王思聪"
	if s == "王思聪" { // 不等于这样写 !=
		fmt.Println("左手一个妹子，右手一个大妈~~~")
	}

	// 30_if支持初始化语句
	//if支持一个初始化语句，初始化语句和判断条件以分号隔开，如下：
	if a := 999; a == 999 {
		fmt.Println("a 是否等于 999呢？答案是yes！")
	}
}
