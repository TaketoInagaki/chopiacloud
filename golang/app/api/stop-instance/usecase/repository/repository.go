package repository

import "app/model"

// ServiceRepository はデータ操作のインターフェースです。
type ServiceRepository interface {
	GetInstance(id uint32) (model.Instance, error)
	StopInstance(id uint32) error
}
