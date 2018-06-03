/*
* @Author: tkk1234
* @Date:   2018-06-03 11:49:09
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-03 14:14:12
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	haha := "abc"
	for i := 0; i < len(haha); i++ { //用for来打印abc中的每个字符
		fmt.Printf("haha[%d]=%c\n", i, haha[i])
	}
	//更好的写法可以写成 迭代的方式
	//迭代打印每个元素，默认返回2个值：一个是元素位置，一个是元素本身。结果同上
	for i, data := range haha {
		fmt.Printf("haha[%d]=%c\n", i, data)
	}

	for i := range haha { //第2个返回值默认丢弃，返回元素的位置（下标）。结果同上
		fmt.Printf("haha[%d]=%c\n", i, haha[i])
	}
	//迭代方式用的很多，配合切片，通道channel等方式
}
