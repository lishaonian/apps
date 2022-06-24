package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)

}

func post() {
	resp, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body.Close() }()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)

}
