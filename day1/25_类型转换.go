/*
* @Author: tkk1234
* @Date:   2018-05-30 22:23:13
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-30 23:32:54
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	var (
		flag bool = true
		ch   byte = 'a'     //字符型本质上就是整形
		t         = int(ch) //用int将a字符转换成int格式，因为字符型本质上就是整形，所以能互转。
	)
	fmt.Printf("flag = %t\n", flag) //%t以true或false输出布尔值
	//bool类型不能转换为int，int也不能转换为bool
	fmt.Printf("flage = %d\n", flag) //%d 输出为int整形值。所以此行结果输出不正确。
	//
	fmt.Println("t = ", t)
}
