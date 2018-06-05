/*
* @Author: tkk1234
* @Date:   2018-06-04 21:19:46
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-04 22:46:51
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func ok(args ...int) { //传递实参可以使0个或多个
	fmt.Println("len(args) = ", len(args)) //获取用户传递参数的个数
	fmt.Println("======================")

	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data) //返回2个值（data名字自取），第一个是下标，第二个是下标所对应的数
	}
}

func main() {
	ok()
	fmt.Println("======================")
	ok(22)
	fmt.Println("======================")
	ok(22, 33, 55)
}

//固定参数（a int）一定要给它传参，
//后面的不定参数（args...int）根据需要传参，而且这2者顺序不能颠倒着写。
func kk(a int, args ...int) {

}

func main() {
	kk(55, 222)
}
