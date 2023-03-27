package model

// SignatureResponse 签名请求返回的结果
type SignatureResponse struct {
	CdnUrl       string `json:"cdnUrl"`
	ResourcePath string `json:"resourcePath"`
	UploadUrl    string `json:"uploadUrl"`
}

// SignaturePayLoad 发送拿取签名请求时的请求体
type SignaturePayLoad struct {
	Method      string `json:"method"`
	Resource    string `json:"resource"`
	ContentType string `json:"contentType"`
}
