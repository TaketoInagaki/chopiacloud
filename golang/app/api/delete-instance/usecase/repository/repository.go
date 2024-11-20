package repository

import "app/model"

// ServiceRepository はデータ操作のインターフェースです。
type ServiceRepository interface {
	GetInstance(id uint32) (model.Instance, error)
	DeleteInstance(id uint32) error
}
