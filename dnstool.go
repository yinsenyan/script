//需求：
//0. 通过参数指定集群、机器、线路
//1. 停止某急群中某台服务器的IP的解析
//2. 停止某集群中某个运营商的IP的解析
//3. 如果2中停止的是默认运营商，则选定另一个运营商修改为默认
//4. 恢复到预期的解析状态
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

type response struct {
	Status status   `json:"status"`
	Domain domain   `json:"domain"`
	Info   string   `json:"info"`
	Record []record `json:"records"`
}

type status struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Created_at string `json:"created_at"`
}

type domain struct {
	Id     float64 `json:"id"`
	Name   string  `json:"name"`
	Owner  string  `json:"owner"`
	Status string  `json:"status"`
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
	var info response
	api := t.Token + "&format=" + t.Format + "&domain_id=" + t.Domain_id //+ "&record_id=" + t.Record_id
	r, err := http.Post(t.Url, "application/x-www-form-urlencoded", strings.NewReader(api))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &info)
	fmt.Println(info.Domain.Status)
}
