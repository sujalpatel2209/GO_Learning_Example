package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {

	file, err := os.Create("./Sample/Sample.txt")

	if err != nil{

		log.Fatal(err)
	}

	file.WriteString("Sample File Text")

	file.Close()

	data, err := ioutil.ReadFile("./Sample/Sample.txt")

	if err != nil{

		log.Fatal(err)
	}

	fmt.Println("File Data : ", string(data))
}
