package request

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"pic2leetcode/model"
	"strings"
)

func GetSignature(url string, cookie string) (model.SignatureResponse, error) {
	var resp model.SignatureResponse

	// 组装请求
	Payload := model.SignaturePayLoad{
		Method:      "PUT",
		Resource:    "image.png",
		ContentType: "image/png",
	}
	js, err := json.Marshal(Payload)
	if err != nil {
		log.Println("json marshal err : ", err)
		return resp, err
	}

	parms := ioutil.NopCloser(strings.NewReader(string(js)))

	request, err := http.NewRequest("POST", url, parms)
	if err != nil {
		log.Println("new request err : ", err)
		return resp, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Cookie", cookie)

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		log.Println("send request err : ", err)
		return resp, err
	}

	// 读取 body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
		panic(err)
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	return resp, nil
}
