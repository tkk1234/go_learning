/*
* @Author: tkk1234
* @Date:   2018-06-03 14:45:49
* @Last Modified by:   tkk1234
* @Last Modified time: 2018-06-03 15:25:45
 */
package main // 必须有一个main包（和下面的mian函数是2个不同的东西）

import (
	"fmt"
)

func main() {
	fmt.Println("111111")
	goto end //goto可以用在任何地方，但不能跨函数使用，end叫做标签，名字自取
	fmt.Println("222222")
end:
	fmt.Println("333333")
}
