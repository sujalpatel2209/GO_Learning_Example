package main

import "fmt"

func main()  {

	num1, num2 := multiReturn(5)

	fmt.Println(num1, num2)

}

func multiReturn(num int) (int, int)   {

	return num+1, num+2
}