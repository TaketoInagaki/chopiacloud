package repository

import "app/api/upload-key/domain"

// ServiceRepository はデータ操作のインターフェースです。
type ServiceRepository interface {
	CreateKey(param domain.RequestParam) (uint32, error)
}
