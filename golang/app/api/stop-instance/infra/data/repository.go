package data

import (
	"app/api/stop-instance/domain"
	"app/api/stop-instance/usecase/repository"
	"app/enum"
	"app/model"
	"errors"
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

// ホストの空き状況を取得
func (t *serviceRepository) GetInstance(id uint32) (model.Instance, error) {
	var instance model.Instance
	err := t.db.Where("id = ?", id).
		Where("status != ?", enum.InstanceStatus.Deleted.Key()).
		First(&instance).Error

	fmt.Println(instance)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Instance{}, &domain.NotFoundError{}
		}

		return model.Instance{}, err
	}

	return instance, nil
}

// インスタンス状態更新
func (t *serviceRepository) StopInstance(id uint32) error {

	data := &model.Instance{
		Status:  enum.InstanceStatus.NotStarted.Key(),
		Updated: time.Now(),
	}
	fmt.Println(data)
	fmt.Println(id)
	err := t.db.Where("id = ?", id).
		Omit("host_ip", "guest_ip", "name", "instance_path", "created").
		Updates(&data).Error
	if err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}
