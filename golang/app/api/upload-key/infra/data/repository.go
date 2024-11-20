package data

import (
	"app/api/upload-key/domain"
	"app/api/upload-key/usecase/repository"
	"app/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

// NewServiceRepository はServiceRepositoryのインターフェースを返します。
func NewServiceRepository(db *gorm.DB) repository.ServiceRepository {
	return &serviceRepository{db: db}
}

func (t *serviceRepository) Begin() *gorm.DB {
	return t.db.Begin()
}

// インスタンス状態更新
func (t *serviceRepository) CreateKey(param domain.RequestParam) (uint32, error) {

	sshKey := &model.SSHKey{
		PublicKey: param.PublicKey,
		Name:      param.Name,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	err := t.db.Create(sshKey).Error
	if err != nil {
		fmt.Println(err)

		return 0, err
	}

	return sshKey.ID, nil
}
