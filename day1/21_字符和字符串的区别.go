/*
* @Author: tkk1234
* @Date:   2018-05-28 22:04:42
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-28 22:43:18
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	//这次我全用自动推导类型来赋值
	str1 := 'a' //字符（就是ASCII码中的单个字母，所以输出会变为ascii码中的编码97）
	str2 := "a" //字符串

	//字符：
	//1，单引号
	//2，字符，往往只有一个字符，转义字符除外'\n'

	//字符串：
	//1，双引号
	//2，字符串有一个或多个字符组成
	//3，字符串后面都隐藏了一个结束符，\0
	fmt.Println("str1 = ", str1)
	fmt.Println("str2 = ", str2)

	//切片
	welcome := "Hello go"
	fmt.Printf("welcome[0] = %c, welcome[1] = %c\n", welcome[0], welcome[1])
}
