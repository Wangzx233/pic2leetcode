package main

import (
	"awesomeProject/load"
	"awesomeProject/request"
)

func main() {
	var err error
	request.Cookie, err = load.GetCookieFromFile()
	if err != nil {
		return
	}

	load.GetResource("./resource")
}
