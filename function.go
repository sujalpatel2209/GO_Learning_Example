package main

import "fmt"

func main() {

	number := [] float64 {2,7,4,5,6,6}

	fmt.Println("Total", sum(number))

}

func sum(arrayValue [] float64) float64 {

	total := 0.0;

	for _,val:=range arrayValue {

		total += val
	}

	return total

}