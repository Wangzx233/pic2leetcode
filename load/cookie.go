package load

import (
	"io/ioutil"
	"log"
)

// GetCookieFromFile 从文件读取用户 cookie
func GetCookieFromFile() (cookie string, err error) {
	content, err := ioutil.ReadFile("cookie.txt")
	if err != nil {
		log.Fatal(err)
	}
	cookie = string(content)
	return
}
