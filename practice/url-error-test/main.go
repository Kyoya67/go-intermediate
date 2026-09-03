package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	client := new(http.Client)

	_, err := client.Get("fooooo")
	if err != nil {
		e := err.(*url.Error)

		fmt.Println("Op:", e.Op)
		fmt.Println("URL:", e.URL)
		fmt.Println("Err:", e.Err)
		fmt.Println(err)
	}
}
