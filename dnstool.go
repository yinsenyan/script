package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type dnsServer struct {
	url    string
	token  string
	format string
}

type record struct {
	domain      string
	record_id   int
	sub_domain  string
	record_type string
	record_line string
	value       string
	status      string
	format      string
}

func main() {
	url := "https://dnsapi.cn/Record.List"
	fmt.Println("Current domain record list")
	getRecordList(url)
}

func getRecordList(url string) {
	r, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(token))
	if err != nil {
		fmt.Println("Get domain list err")
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
}
