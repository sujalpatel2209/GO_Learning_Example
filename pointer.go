package main

import "fmt"

func main()  {

	x := 1

	changeValue(&x)

	fmt.Println("change value : ", x)
}

func changeValue(x *int){

	*x = 5
}