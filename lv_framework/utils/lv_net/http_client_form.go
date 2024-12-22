package lv_net

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostFormToken(url string, data url.Values, token string) (string, error) {
	//把post表单发送给目标服务器
	client := &http.Client{}
	//生成要访问的url
	request, err := http.NewRequest("POST", url, nil)
	//增加header选项
	request.Header.Add("token", token)
	request.Header.Add("Cookie", "test=123456")
	request.Header.Add("User-Agent", "test")
	request.Header.Add("X-Requested-With", "test")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), nil
}

func PostForm(url string, data url.Values) (string, error) {
	//把post表单发送给目标服务器
	res, err := http.PostForm(url, data)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), nil
}

func PostFormMap(httpUrl string, mp map[string]string) (string, error) {
	//把post表单发送给目标服务器
	params := url.Values{}
	for k, v := range mp {
		params.Add(k, v)
	}
	return PostForm(httpUrl, params)
}

func Get(url string) (string, error) {
	//把post表单发送给目标服务器
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return string(body), nil
}
