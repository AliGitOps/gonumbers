package main

import (
	"fmt"
)

func main() {
	var integer int // 用户传参
	/* 更改开头字体输入颜色
	   输出的颜色为绿色
	*/
	green := "\033[1;32m"
	reset := "\033[0m"

	fmt.Println(green + "=====欢迎来到猜数字游戏=====\n" + reset)
	fmt.Println("请输入一个 0 - 100 的整数: ")
	fmt.Scan(&integer)

	///* 判断用户输入的内容是否是规定内的内容,如果不是将重新执行 */

	//swi(integer)

}

func swi(shuzi int) {
	switch {
	case shuzi == 100:

		fmt.Println("特等奖")

	case shuzi >= 90 && shuzi < 100:

		fmt.Println("一等奖")

	case shuzi >= 80 && shuzi < 90:

		fmt.Println("二等奖")

	default:
		fmt.Println("猜错了,没有奖")

	}
}
