package data

import (
	"app/api/create-instance/domain"
	"app/api/create-instance/usecase/repository"
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
func (t *serviceRepository) GetHostAvailability() ([]domain.HostAvailability, error) {
	var instanceAvailabilities []domain.HostAvailability
	err := t.db.Table("instances").
		Select(`
			host_ip,
			CASE WHEN COUNT(CASE WHEN status IN (1, 2, 3) THEN 1 END) >= 2 THEN false ELSE true END AS is_available,
			COUNT(CASE WHEN status IN (1, 2, 3) THEN 1 END) AS count
		`).
		Where("status != ?", enum.InstanceStatus.Deleted.Key()).
		Group("host_ip").
		Find(&instanceAvailabilities).
		Error
	if err != nil {
		return nil, err
	}

	fmt.Println(instanceAvailabilities)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &domain.NotFoundError{}
		}

		return nil, err
	}

	return instanceAvailabilities, nil
}

// SSHキー取得
func (t *serviceRepository) GetSSHKey(sshKeyId uint32) (model.SSHKey, error) {
	var sshKey model.SSHKey
	err := t.db.Where("id = ?", sshKeyId).First(&sshKey).Error
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SSHKey{}, &domain.NotFoundError{}
		}

		return model.SSHKey{}, err
	}

	return sshKey, nil
}

// インスタンス作成
func (t *serviceRepository) CreateInstance(hostIP string, name string) (uint32, error) {

	Instance := &model.Instance{
		HostIP:  hostIP,
		Status:  enum.InstanceStatus.InOperation.Key(),
		Name:    name,
		Created: time.Now(),
		Updated: time.Now(),
	}
	err := t.db.Create(Instance).Error
	if err != nil {
		fmt.Println(err)

		return 0, err
	}

	return Instance.ID, nil
}
