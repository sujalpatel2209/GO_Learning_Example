package main

import (
	"net/http"
	"fmt"
)

func main()  {

	http.HandleFunc(`/`, func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "Sujal Patel")

	})

	http.ListenAndServe(":2222", nil)
}
