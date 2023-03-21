package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func GetImage(imageURL string) []byte {

	// 获取图片
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Println("获取图片失败：", err)
		fmt.Println(runtime.Caller(1))

		return nil
	}
	defer resp.Body.Close()

	// 将图片转换为字节切片
	imgBytes := new(bytes.Buffer)
	imgBytes.ReadFrom(resp.Body)

	return imgBytes.Bytes()
}
