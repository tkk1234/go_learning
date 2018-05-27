/*
* @Author: iphone5s
* @Date:   2018-05-27 22:25:54
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-05-27 22:54:02
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	var (
		haha byte = 97
	)
	// 入口函数 main（就像一个商场的入口）,大括号必须和函数“main”这个英文同行！
	// Println会自动换行
	// 调用函数大部分都需要调用包，即上面的import fmt格式化包。
	fmt.Printf("%c,%d\n", haha, haha) //%c以字符打印，%d以整形方式打印
	//得出结果a，97 就表示a在内存中的数字为97，验证了ASCII码表中的映射
	//再测试一下
	haha = 'a' //字符要用单引号，将a赋值给haha
	fmt.Printf("%c,%d\n", haha, haha)

	// 大写转小写，小写转大写，大小写相差32，小写大结果为65，大写小结果为97
	fmt.Printf("大写A ascii码：%d,小写a ascii码%d", 'A', 'a')
	fmt.Printf()
}
