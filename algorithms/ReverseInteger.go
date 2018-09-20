package main

import "fmt"

//反转整数

func reverse(x int) int {
	var result int
	for x != 0 {
		//取出最右的数字
		r := x % 10
		x /= 10
		if result > 2147483647 / 10 || (result == 2147483647 / 10 && r > 7) {
			return 0
		}
		if result < -2147483648 /10 || (result == -2147483648 / 10 && r < -8) {
			return 0
		}
		result = result * 10 + r;
	}
	return result
}

func main() {
	fmt.Println(reverse(12345))
}
