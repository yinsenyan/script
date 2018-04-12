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
	"flag"
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
	Id          string `json:"id"`
	Value       string `json:"value"`
	Enabled     string `json:"enabled"`
	Status      string `json:"status`
	Name        string `json:"name"`
	Line        string `json:"line"`
	Record_type string `json:"type"`
	Remark      string `json:"remark"`
}

type api struct {
	Rl        string `json:"recordlist"`
	Rm        string `json:"recordmodify"`
	Rs        string `json:"recordstatus"`
	Token     string `json:"token"`
	Format    string `json:"format"`
	Domain_id string `json:"domain_id"`
}

var h = flag.Bool("h", false, "print help info")
var help = flag.Bool("help", false, "print help info")
var isp = flag.String("isp", "", "define isp, such:电信、联通、移动")
var clu = flag.String("cluster", "", "define cluster, such:nb、xs、xq")
var sta = flag.String("status", "", "set record status, such: enable、disable")

func unmarshalApi() (a api) {
	apifile, err := ioutil.ReadFile("api.json")
	if err != nil {
		fmt.Println("error open file api.json")
	}
	json.Unmarshal(apifile, &a)
	return
}

func changeDefaultRecord(cluster, record_status string, res allinfo, a api) (response allinfo) {
	if record_status == "disable" {
		for i := 0; i < len(res.Record)-1; i++ {
			iscluster, _ := regexp.MatchString("ke-"+cluster, res.Record[i].Name)
			if res.Record[i].Line == "联通" && iscluster {
				postdate := a.Token + "&format=" + a.Format + "&domain_id=" + a.Domain_id + "&record_id=" + res.Record[i].Id + "&sub_domain=" + res.Record[i].Name + "&value=" + res.Record[i].Value + "&record_type=A" + "&record_line=默认"
				response = invocateAPI(a.Rm, postdate)
			}
		}
	}
	if record_status == "enable" {
		for i := 0; i < len(res.Record)-1; i++ {
			iscluster, _ := regexp.MatchString(cluster, res.Record[i].Name)
			if res.Record[i].Line == "默认" && iscluster {
				postdate := a.Token + "&format=" + a.Format + "&domain_id=" + a.Domain_id + "&record_id=" + res.Record[i].Id + "&sub_domain=" + res.Record[i].Name + "&value=" + res.Record[i].Value + "&record_type=A" + "&record_line=" + res.Record[i].Remark
				response = invocateAPI(a.Rs, postdate)
			}
		}
	}
	return
}

func changeRecordStatus(cluster, isp, record_status string, res allinfo, a api) (response allinfo) {
	for i := 0; i < len(res.Record)-1; i++ {
		iscluster, _ := regexp.MatchString("ke-"+cluster, res.Record[i].Name)
		isisp, _ := regexp.MatchString(isp, res.Record[i].Line)
		if iscluster && isisp {
			postdate := a.Token + "&format=" + a.Format + "&domain_id=" + a.Domain_id + "&record_id=" + res.Record[i].Id + "&status=" + record_status
			response = invocateAPI(a.Rs, postdate)
		}
	}
	return
}

func getRecordList(a api) (response allinfo) {
	postdate := a.Token + "&format=" + a.Format + "&domain_id=" + a.Domain_id
	response = invocateAPI(a.Rl, postdate)
	return
}

func invocateAPI(url, postdate string) (response allinfo) {
	r, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postdate))
	if err != nil {
		fmt.Println("http request err : ", err.Error())
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("get domain response body error")
	}
	json.Unmarshal(body, &response)
	// fmt.Println(response.Status, response.Domain, response.Info)
	// fmt.Println(postdate)
	return
}

func main() {
	a := unmarshalApi()
	res := getRecordList(a)

	flag.Parse()
	//The help info will be output when argument is -h/-help/--help, and list the all record
	if *help || *h {
		for i := 0; i <= len(res.Record)-1; i++ {
			fmt.Println(res.Record[i].Name, res.Record[i].Line, res.Record[i].Value, res.Record[i].Enabled)
		}
		fmt.Println("输入参数，例如：--cluster=nb --isp=联通 --status=disable")
	} else {
		//If disable a default record , auto set the default is cucc
		if *isp == "默认" && *sta == "disable" {
			fmt.Println("You will disable the default record, the cucc will be set default record")
			r2 := changeDefaultRecord(*clu, *sta, res, a)
			fmt.Println("Set cucc to default, Response:", r2.Status.Message)
		}
		//if enable a default record, auto set the current default record to real line
		if *isp == "默认" && *sta == "enable" {
			fmt.Println("Recover a default record, the current default will be set real line(in record.remart)")
			r2 := changeDefaultRecord(*clu, *sta, res, a)
			fmt.Println("Recover real default, Response:", r2.Status.Message)
		}
		r1 := changeRecordStatus(*clu, *isp, *sta, res, a)
		fmt.Println("Disable some record, Response:", r1.Status.Message)
	}
}
