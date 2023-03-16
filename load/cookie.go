package load

import (
	"io/ioutil"
	"log"
)

func GetCookieFromFile() (cookie string, err error) {
	content, err := ioutil.ReadFile("cookie.txt")
	if err != nil {
		log.Fatal(err)
	}
	cookie = string(content)
	return
}
