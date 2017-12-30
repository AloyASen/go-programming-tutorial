package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.nerve.org.in/")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	resp.Body.Close()
	fmt.Println(string_body)
}
