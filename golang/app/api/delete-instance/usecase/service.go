package usecase

import (
	"app/api/delete-instance/usecase/repository"
	"app/enum"
	"fmt"

	"os"
	"os/exec"
)

type service struct {
	rep repository.ServiceRepository
}

// Service はインターフェースです。
type Service interface {
	DeleteInstance(uint32) error
}

// NewService は初期化済みのServiceを返す。
func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// インスタンス作成
func (s *service) DeleteInstance(id uint32) error {
	// インスタンス状態の取得
	instance, err := s.rep.GetInstance(id)
	if err != nil {
		return err
	}

	//instanceのstatusが0以外の場合はエラー
	if instance.Status == enum.InstanceStatus.Running.Key() || instance.Status == enum.InstanceStatus.InOperation.Key() {
		return fmt.Errorf("please stop the instance first")
	}

	vmUser := os.Getenv("VM_USER")
	sshKeyPath := os.Getenv("SSH_KEY_PATH")
	// 同じPC内に存在する仮想マシン上で動かしているDockerで既存のコンテナを削除する
	cmd := exec.Command("ssh", "-o", "StrictHostKeyChecking=no", "-i", sshKeyPath, vmUser+"@"+instance.HostIP, "docker rm -f", instance.Name)
	fmt.Println(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return err
	}
	// インスタンスの状態を確認
	cmd = exec.Command("ssh", "-o", "StrictHostKeyChecking=no", "-i", sshKeyPath, vmUser+"@"+instance.HostIP, "docker ps -a")
	fmt.Println(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return err
	}

	// インスタンスの空きを更新（固定値ステータス）
	err = s.rep.DeleteInstance(id)
	if err != nil {
		return err
	}

	fmt.Println("Instance deleted")
	return nil
}
