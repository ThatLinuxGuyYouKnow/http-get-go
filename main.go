package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, error := http.Get("https://jsonplaceholder.typicode.com/posts")
	if error == nil { // 'nil' not null as in Dart and Js

		body, error := io.ReadAll(response.Body)
		sb := string(body)
		if error == nil {
			println(sb)
		}
	}
	fmt.Println("error", error)
}
