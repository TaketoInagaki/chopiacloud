package domain

type RequestParam struct {
	InstanceID uint32 `json:"instance_id"`
}

// 該当データなしエラー格納用
type NotFoundError struct {
	error
}
