package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func HandlerRequest(method string,url string,headerMap map[string]string,postBody map[string]interface{}) ([]byte,error) {

	data, errMarshal := json.Marshal(postBody)
	if errMarshal != nil {
		return nil,errMarshal
	}
	r, _ := http.NewRequest(method, url,strings.NewReader(string(data)))
	// 循环添加请求头
	for k,v:= range headerMap {
		r.Header.Add(k,v)
	}
	resp ,err := http.DefaultClient.Do(r)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	return body,nil
}