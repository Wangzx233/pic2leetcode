package main

import (
	"pic2leetcode/load"
	"pic2leetcode/request"
)

func main() {
	var err error
	request.Cookie, err = load.GetCookieFromFile()
	if err != nil {
		return
	}

	load.GetResource("./resource")
}
