package lv_net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var httpClient *http.Client

func Init() {
	log.Println("----------初始化httpUtils---------------")
	httpClient = &http.Client{Transport: http.DefaultTransport}
	httpClient.Timeout = 5 * time.Second
}

// 200 成功 其它 失败
func PostJSON(url string, data interface{}) error {
	bytesData, _ := json.Marshal(data)
	//global.Logbiz.Infof("====>发送url[%s] 数据[%v] ", url, string(bytesData))
	resp, err := httpClient.Post(url, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		//global.Logbiz.Errorf("发送异常[%s] panic:[%v]", url, err)
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//jsonStr := string(body)
	//global.Logbiz.Infof("<=====发送完成: [%v] ", jsonStr)
	var datamap map[string]interface{}
	err = json.Unmarshal(body, &datamap)
	if err == nil {
		return err
	} //
	code := datamap["code"].(string)
	if code == "200" {
		return nil
	} else {
		msg := datamap["msg"].(string)
		return errors.New(msg)
	}

}

func PostJsonAndHeader(url string, json []byte, headers map[string]string) ([]byte, error) {
	//把post表单发送给目标服务器
	client := &http.Client{}
	//生成要访问的url
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	request.Header.Set("Content-Type", "application/json")
	//增加header选项
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	request.Header.Add("X-Requested-With", "test")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	return body, nil
}
