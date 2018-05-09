package main

import "fmt"

func main() {

	var a int = 10;
	var b int = 20;

	var t = true;

	str := "String Test";

	arrayVal := [5] int {2,3,4,5,6}

	for _, val := range arrayVal {

		fmt.Println(val);
	}


	fmt.Println(a+b)

	fmt.Println(str)

	fmt.Println(!t);
}