package sotest

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

var i = 0

func Inc() int {
	i = i + 1
	return i
}

type HttpRequest struct {
	StrArray []string         `json:"strArray"`
	StrMap   map[string]int64 `json:"strMap"`
	Str      string           `json:"Str"`
	intValue int64            `json:"intValue"`
}

type HttpResponse struct {
	Str      string           `json:"Str"`
	StrArray []string         `json:"strArray"`
	StrMap   map[string]int64 `json:"strMap"`
}

func Add(x, y int64) int64 {
	return x + y
}

func ToUpperCase(str string) string {
	return strings.ToUpper(str)
}

func StructInvoke(req *HttpRequest) *HttpResponse {
	resp := &HttpResponse{}
	resp.Str = req.Str
	arr := make([]string, len(req.StrArray))
	makeUp := make(map[string]int64)
	for i, s := range req.StrArray {
		arr[i] = s
	}
	resp.StrArray = arr
	for s, i := range req.StrMap {
		makeUp[s] = i
	}
	resp.StrMap = makeUp
	return resp
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
