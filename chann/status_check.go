package main

import (
	"fmt"
	"net/http"
)

func Creating_links() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazom.com",
	}
	// In this we will receive the data from the channels
	c := make(chan string)
	for _, links := range links {
		go CheckLink(links, c)
	}
	// After receiving the msg we are just printing the message received from the channels
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

// This is about how we are communicating betweent the links and channels
// In this we sent the data from the channels mean child channels
func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down! :")
		// We are passing the message to the channels
		c <- "might be down i think"
		return
	}
	fmt.Println(link)
	// Also we are passing the message to the channels
	c <- "Yes it is up"

}
func main() {
	Creating_links()
}
