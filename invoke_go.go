package sotest

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func Add(x, y int64) int64 {
	return x + y
}

func ToUpperCase(str string) string {
	return strings.ToUpper(str)
}

func ReqThird(reqPath string, url string) string {
	filepath.Join()
	Client := resty.New()
	fmt.Println("-----------------入参:", reqPath)
	post, err := Client.R().Post(url)
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	if err != nil {
		str := err.Error()
		return fmt.Sprintf(`{"code":400,"msg":%s}`, str)
	}
	return string(post.Body())
}

func ReqOrig(reqPath string, url string) string {

	req, _ := http.NewRequest("POST", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
