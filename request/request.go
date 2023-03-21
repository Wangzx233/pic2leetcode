package request

import "fmt"

var Cookie string

// UpImageAndGetUrl 将request操作汇总,拿取图片然后发送，然后拿到转换后的url
func UpImageAndGetUrl(imgUrl string) (CdnUrl string, err error) {
	// 获取图片字节流
	imageBytes := GetImage(imgUrl)

	fmt.Println("下载图片成功")

	signature, err := GetSignature("https://oss.leetcode.cn/signature", Cookie)
	if err != nil {
		return signature.CdnUrl, err
	}

	fmt.Println("获取签名成功")

	err = UpImage(signature, imageBytes)
	if err != nil {
		return signature.CdnUrl, err
	}
	fmt.Println("上传图片成功")

	return signature.CdnUrl, err
}
