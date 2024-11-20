package repository

import (
	"app/api/create-instance/domain"
	"app/model"
)

// ServiceRepository はデータ操作のインターフェースです。
type ServiceRepository interface {
	GetHostAvailability() ([]domain.HostAvailability, error)
	GetSSHKey(sshKeyId uint32) (model.SSHKey, error)
	CreateInstance(hostIP string, name string) (uint32, error)
}
