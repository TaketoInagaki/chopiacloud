package usecase

import (
	"app/api/upload-key/domain"
	"app/api/upload-key/usecase/repository"
	"fmt"
)

type service struct {
	rep repository.ServiceRepository
}

// Service はインターフェースです。
type Service interface {
	UploadKey(domain.RequestParam) (uint32, error)
}

// NewService は初期化済みのServiceを返す。
func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// インスタンス作成
func (s *service) UploadKey(param domain.RequestParam) (uint32, error) {
	// sshKeyを登録
	id, err := s.rep.CreateKey(param)
	if err != nil {
		return 0, err
	}

	fmt.Println("Instance stopped")
	return id, nil
}
