package lib_net

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetParam(c *gin.Context, key string) string {
	v := c.Query(key)
	if v == "" {
		v = c.PostForm(key)
		if v == "" {
			v = c.GetHeader(key)
		}
	}
	return v
}
func PostFormAndHeader(url string, data url.Values, token string) (string, error) {
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

func Get(url string) (string, error) {
	//把post表单发送给目标服务器
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), nil
}
