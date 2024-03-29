package nekosrun

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func http_get_param_serialize(params map[string]interface{}) string {
	param_url := ""
	for param, obj := range params {
		if op, ok := obj.(string); ok {
			param_url += param + "=" + url.QueryEscape(op) + "&"
		}
		if op, ok := obj.(int); ok {
			param_url += param + "=" + strconv.Itoa(op) + "&"
		}
		if op, ok := obj.(bool); ok {
			if op {
				param_url += param + "=1&"

			} else {
				param_url += param + "=0&"

			}
		}
		if op, ok := obj.(float32); ok {
			param_url += param + "=" + fmt.Sprintf("%f", op) + "&"
		}
		if op, ok := obj.(float64); ok {
			param_url += param + "=" + fmt.Sprintf("%f", op) + "&"
		}
	}
	param_url = param_url[:len(param_url)-1]
	return param_url
}

func http_get(url string, params *map[string]interface{}) []uint8 {
	client := &http.Client{}
	if params != nil {
		url += "?" + http_get_param_serialize(*params)
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body
}

func http_post(url string, data string, cookies string) []uint8 {
	client := &http.Client{}

	reqBody := strings.NewReader(data)
	request, _ := http.NewRequest("POST", url, reqBody)

	if len(cookies) > 0 {
		request.Header.Add("Cookie", cookies)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 11; M2002J9E Build/RKQ1.200826.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/77.0.3865.120 MQQBrowser/6.2 TBS/045525 Mobile Safari/537.36 MMWEBID/2919 MicroMessenger/8.0.2.1860(0x2800023B) Process/tools WeChat/arm64 Weixin NetType/5G Language/zh_CN ABI/arm64")
	// request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return body
}
