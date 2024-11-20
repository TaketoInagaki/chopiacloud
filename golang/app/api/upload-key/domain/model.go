package domain

type RequestParam struct {
	PublicKey string `json:"public_key"`
	Name      string `json:"name"`
}

// 該当データなしエラー格納用
type NotFoundError struct {
	error
}
