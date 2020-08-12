package main

import "fmt"


func main(){
	result := cal(12, 12, '+')
	fmt.Println("result =", result)
}

func cal(n1 float64, n2 float64, operator byte) float64{
	var res float64
	switch operator{
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误。。。")
	}
	// fmt.Println("res = ", res)
	return res
}