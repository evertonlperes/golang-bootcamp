package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// creating byte slice required by Reader class
	// and setting the elements size that we want to
	// Learning way to do it
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// One line solution
	io.Copy(os.Stdout, resp.Body)

}
