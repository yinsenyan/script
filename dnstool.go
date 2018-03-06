package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type post struct {
	url       string
	token     string
	format    string
	domain_id string
	record_id string
}

type domain struct {
	Status status   `json:"status"`
	Domain string   `json:"domain"`
	Info   string   `json:"info"`
	Record []record `json:"records"`
}

type status struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Created_at string `json:"created_at"`
}

type record struct {
	Id          string `json:“id”`
	Ttl         string `json:"ttl"`
	Value       string `json:"value"`
	Status      string `json:"status`
	Updated_on  string `json:"updated_on"`
	Name        string `json:"name"`
	Line        string `json:"line"`
	Record_type string `json:"type"`
}

type token struct {
	Url       string `json:"url"`
	Token     string `json:"token"`
	Format    string `json:"format"`
	Domain_id string `json:"domain_id"`
	Record_id string `json:"record_id"`
}

func main() {
	t := unmarshalToken()
	getRecordList(t)
}

func unmarshalToken() (t token) {
	tokenFile, err := ioutil.ReadFile("token.json")
	if err != nil {
		fmt.Println("error open file token.json")
	}
	json.Unmarshal(tokenFile, &t)
	return
}

func getRecordList(t token) {
	var d domain
	r, err := http.Post(t.Url, "application/x-www-form-urlencoded", strings.NewReader(t.Token+"&"+t.Format+"&"+t.Domain_id+"&"+t.Record_id))
	if err != nil {
		fmt.Println("Get domain list err")
		fmt.Println(err.Error())
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &d)
	fmt.Println(d)
}
