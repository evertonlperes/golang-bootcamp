package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create the new channel
	c := make(chan string)

	for _, link := range links {
		// Start a new execution in a new
		// go routine
		// 'go' keyword creates a new thread
		// of Go Routine
		go checkLink(link, c)
	}

	// Alternative loop for waiting
	// the channel to return the value
	// and run the checkLink()
	for l := range c {
		go func(link string) { // Function literals
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// just consider the error msg
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link // send the string to the channel
}
