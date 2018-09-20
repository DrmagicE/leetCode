package main

import (
	"fmt"
)



//插入顺序，从第一行递增到第n行，然后从n又回到第一行
//时间复杂度O(n)
func convert(s string, numRows int) string {
	var z [][]byte
	if numRows == 1 {
		return s
	}
	strlen := len(s)
	z = make([][]byte, numRows)
	for i:= 0; i < numRows ;i++ {
		z[i] = make([]byte,0)
	}
	row := 0
	currentFlag := "+"
	for i := 0; i < strlen; i++ {
		z[row] = append(z[row], byte(s[i]))
		if currentFlag == "+" && row == numRows - 1 {
			currentFlag = "-"
		}
		if currentFlag == "-" && row == 0 {
			currentFlag = "+"
		}
		if currentFlag == "+" {
			row++
		} else {
			row--
		}
	}
	var result []byte
	result = make([]byte,0,strlen)
	for i :=  0; i < numRows; i++ {
		result = append(result, z[i]...)
	}
	return string(result)
}

func main() {
	// PA YPA L ISHIRING
	fmt.Println(convert("PAYPALISHIRING",4))
}
