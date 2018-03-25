//需求：
//0. 通过参数指定集群、机器、线路
//1. 停止某集群中某个运营商的IP解析
//2. 恢复某急群众某个运营商的IP解析
//3. 如果2中停止的是默认运营商，则选定另一个运营商修改为默认
//4. 恢复到预期的解析状态

//API列表
//记录列表:https://dnsapi.cn/Record.List
//	curl -X POST https://dnsapi.cn/Record.List -d 'login_token=LOGIN_TOKEN&format=json&domain_id=2317346'
//修改记录:https://dnsapi.cn/Record.Modify
//	curl -X POST https://dnsapi.cn/Record.Modify -d 'login_token=LOGIN_TOKEN&format=json&domain_id=2317346&record_id=16894439&sub_domain=www&value=3.2.2.2&record_type=A&record_line_id=10%3D3'
//修改状态:https://dnsapi.cn/Record.Status
//	curl -X POST https://dnsapi.cn/Record.Status -d 'login_token=LOGIN_TOKEN&format=json&domain_id=2317346&record_id=16894439&status=disable'

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type post struct {
	url       string
	token     string
	format    string
	domain_id string
	record_id string
}

type allinfo struct {
	Status status   `json:"status"`
	Domain domain   `json:"domain"`
	Info   info     `json:"info"`
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

type info struct {
	Sub_domains  string `json:"sub_domains"`
	Record_total string `json:"record_total"`
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

type api struct {
	Rl        string `json:"recordlist"`
	Rm        string `json:"recordmodify"`
	Rs        string `json:"recordstatus"`
	Token     string `json:"token"`
	Format    string `json:"format"`
	Domain_id string `json:"domain_id"`
}

func unmarshalApi() (a api) {
	apifile, err := ioutil.ReadFile("api.json")
	if err != nil {
		fmt.Println("error open file api.json")
	}
	json.Unmarshal(apifile, &a)
	return
}

func getRecordList(a api) (response allinfo) {
	postdate := a.Token + "&format=" + a.Format + "&domain_id=" + a.Domain_id //+ "&record_id=" + a.Record_id
	r, err := http.Post(a.Rl, "application/x-www-form-urlencoded", strings.NewReader(postdate))
	if err != nil {
		fmt.Println(a.Rl)
		fmt.Println("http request err : ", err.Error())
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("get domain response body error")
	}
	json.Unmarshal(body, &response)
	return
}

func main() {
	a := unmarshalApi()
	res := getRecordList(a)
	fmt.Println(res.Record)
	//changeRecordStatus("test", "test", "disable", res, a)
}

func changeRecordStatus(cluster, isp, record_status string, res allinfo, a api) (response allinfo) {
	for i := 0; i < len(res.Record); i++ {
		if b, _ := regexp.MatchString(cluster, res.Record[i].Name); b {
			postdate := a.Token + "&format=" + a.Format + "&domain_id=" + a.Domain_id + "&record_id=" + res.Record[i].Id + "&status=" + record_status
			r, err := http.Post(a.Rs, "application/x-www-form-urlencoded", strings.NewReader(postdate))
			if err != nil {
				fmt.Println("http request err : ", err.Error())
			}
			defer r.Body.Close()
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println("get domain response body error")
			}
			json.Unmarshal(body, &response)
			fmt.Println(response.Status)
		}
	}
	return
}

// func (*record) modifyRecord(api string) {

// }
