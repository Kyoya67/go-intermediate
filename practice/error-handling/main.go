package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	// client := new(http.Client)

	// _, err := client.Get("fooooo")
	// clientErr := err.(*url.Error)

	// fmt.Println("Op:", clientErr.Op)
	// fmt.Println("URL:", clientErr.URL)
	// fmt.Println("Err:", clientErr.Err)
	// fmt.Println(clientErr)

	// fmt.Println()

	// _, err = strconv.Atoi("a")
	// strconvErr := err.(*strconv.NumError)
	// fmt.Println("Func:", strconvErr.Func)
	// fmt.Println("Num:", strconvErr.Num)
	// fmt.Println("Err:", strconvErr.Err)
	// fmt.Println(strconvErr)

	if err := doSomething(); err != nil {
		fmt.Println("Error:", err)
	}
}

func getStrconvNumError() error {
	_, err := strconv.Atoi("a")
	if err != nil {
		return err.(*strconv.NumError)
	}
	return nil
}

func getURLError() error {
	client := new(http.Client)
	_, err := client.Get("fooooo")
	if err != nil {
		return err.(*url.Error)
	}
	return nil
}

func doSomething() error {
	// 処理 1 *strconv.NumError 型を発生させ、返す関数 (自作)
	err := getStrconvNumError()
	if err != nil {
		return err
	}
	// 処理 2 *url.Error 型を発生させ、返す関数 (自作)
	err = getURLError()
	if err != nil {
		return err
	}
	return nil
}
