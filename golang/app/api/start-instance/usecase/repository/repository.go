package repository

import "app/model"

// ServiceRepository はデータ操作のインターフェースです。
type ServiceRepository interface {
	GetInstance(id uint32) (model.Instance, error)
	StartInstance(id uint32) error
}
