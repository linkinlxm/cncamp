// 课后练习 1.1
// 给定一个字符串数组
// [“I”,“am”,“stupid”,“and”,“weak”]
// 用 for 循环遍历该数组并修改为
// [“I”,“am”,“smart”,“and”,“strong”]
package main

import (
	"fmt"
)

func main() {
	array := [5]string{"i", "am", "stupid", "and", "week"}
	for i, _ := range array {
		fmt.Println(array[i])
		switch array[i] {
		case "stupid":
			/* code */
			array[i] = "smart"
		case "week":
			array[i] = "strong"
		}
	}
	fmt.Println("New Array is", array)
	// for _, value := range array{
	// 	fmt.Println(value)
	// }

}
