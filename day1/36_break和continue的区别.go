/*
* @Author: tkk1234
* @Date:   2018-06-03 14:15:11
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-03 14:38:31
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for { //for如果后面不写任何东西，这个死循环永远为真，就进入了死循环。
		i++
		time.Sleep(time.Second)
		if i == 5 {
			break //跳出循环，如果嵌套多个，则跳出最近的那个
			//continue  //跳出本次循环，再继续
		}
		fmt.Println("i = ", i)
	}

}
