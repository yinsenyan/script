package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	word, err := getWord()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		translate(word)
	}
}

type response struct {
	From string   `json:"from"`
	To   string   `json:"to"`
	R    []result `json:"trans_result"`
}

type result struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func getWord() (lang []string, err error) {
	if len(os.Args) != 2 {
		e := fmt.Errorf("Please input a word to translate!")
		return lang, e
	} else {
		word := os.Args[1]
		lang := append(lang, word)
		isok, _ := regexp.MatchString("[a-z]+", word)
		if isok {
			lang = append(lang, "en", "zh")
			return lang, nil
		} else {
			lang = append(lang, "zh", "en")
			return lang, nil
		}
	}
}

func translate(word []string) {
	var r response
	appid := "id"
	password := "password"
	salt := strconv.Itoa(rand.Int())

	str := appid + word[0] + salt + password
	m := md5.New()
	io.WriteString(m, str)
	sign := fmt.Sprintf("%x", m.Sum(nil))
	url := "/api/trans/vip/translate" + "?appid=" + appid + "&q=" + word[0] + "&from=" + word[1] + "&to=" + word[2] + "&salt=" + salt + "&sign=" + sign
	api := "http://api.fanyi.baidu.com" + url

	res, err := http.Get(api)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("error of get api")
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error of read response body")
	}
	json.Unmarshal(content, &r)
	l := len(r.R)
	fmt.Println(r.R[l-1].Src, "--->", r.R[l-1].Dst)
}
