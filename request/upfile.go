package request

import (
	"awesomeProject/model"
	"bytes"
	"log"
	"net/http"
)

func UpImage(signature model.SignatureResponse, imageBytes []byte) error {

	body := bytes.NewBuffer(imageBytes)

	request, err := http.NewRequest("PUT", signature.UploadUrl, body)
	if err != nil {
		log.Println("new request err : ", err)
		return err
	}

	request.Header.Set("Content-Type", "image/png")

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		log.Println("send request err : ", err)
		return err
	}

	return nil
}
