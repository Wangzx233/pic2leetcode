package model

type SignatureResponse struct {
	CdnUrl       string `json:"cdnUrl"`
	ResourcePath string `json:"resourcePath"`
	UploadUrl    string `json:"uploadUrl"`
}

type SignaturePayLoad struct {
	Method      string `json:"method"`
	Resource    string `json:"resource"`
	ContentType string `json:"contentType"`
}
