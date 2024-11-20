package usecase

import (
	"app/api/create-instance/domain"
	"app/api/create-instance/usecase/repository"
	"fmt"
	"path/filepath"
	"strconv"

	"os"
	"os/exec"
)

type service struct {
	rep repository.ServiceRepository
}

// Service はインターフェースです。
type Service interface {
	CreateInstance(uint32) (domain.Response, error)
}

// NewService は初期化済みのServiceを返す。
func NewService(r repository.ServiceRepository) Service {
	return &service{rep: r}
}

// インスタンス作成
func (s *service) CreateInstance(sshKeyId uint32) (domain.Response, error) {
	// インスタンスの空きを確認
	hosts, err := s.rep.GetHostAvailability()
	if err != nil {
		return domain.Response{}, err
	}
	var hostIP string
	var containerName string
	// 利用可能なhostがない場合
	allUnavailable := true
	for _, host := range hosts {
		if host.Available {
			allUnavailable = false
			break
		}
	}
	if len(hosts) == 2 && allUnavailable {
		fmt.Println("No available hosts")
		return domain.Response{}, fmt.Errorf("no available hosts")
	}

	// 利用可能なhostを取得
	for _, host := range hosts {
		if host.Available {
			hostIP = host.HostIP
			containerName = hostIP + "-" + strconv.Itoa(host.Count+1)
			break
		}
	}

	// SSH接続情報の取得
	if hostIP == "" {
		for _, host := range hosts {
			if host.HostIP == os.Getenv("VM_IP") && !host.Available {
				hostIP = os.Getenv("VM2_IP")
				break
			}
		}
		if hostIP == "" {
			hostIP = os.Getenv("VM_IP")
		}
		containerName = hostIP + "-1"
	}
	vmUser := os.Getenv("VM_USER")
	sshKeyPath := os.Getenv("SSH_KEY_PATH")
	defaultImage := os.Getenv("DEFAULT_IMAGE")
	fmt.Println(defaultImage)

	// データベースからSSH公開鍵を取得
	ssh, err := s.rep.GetSSHKey(sshKeyId)
	if err != nil {
		return domain.Response{}, err
	}

	// 仮想マシンの .ssh ディレクトリに公開鍵を保存
	authorizedKeyPath := filepath.Join("/home", vmUser, ".ssh", containerName, "authorized_keys")

	// Create the directory for authorized keys
	cmd := exec.Command("ssh", "-o", "StrictHostKeyChecking=no", "-i", sshKeyPath, vmUser+"@"+hostIP,
		fmt.Sprintf("mkdir -p %s", filepath.Dir(authorizedKeyPath)))
	fmt.Println(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return domain.Response{}, err
	}

	// Save the public key to the authorized keys file
	cmd = exec.Command("ssh", "-o", "StrictHostKeyChecking=no", "-i", sshKeyPath, vmUser+"@"+hostIP,
		fmt.Sprintf("echo \"%s\" >> %s", ssh.PublicKey, authorizedKeyPath))
	fmt.Println(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return domain.Response{}, err
	}

	// Dockerコンテナを起動し、ポートを公開し、公開鍵をマウント
	cmd = exec.Command("ssh", "-o", "StrictHostKeyChecking=no", "-i", sshKeyPath, vmUser+"@"+hostIP,
		fmt.Sprintf("docker run -d -v %s:/root/.ssh/authorized_keys --name %s %s", authorizedKeyPath, containerName, defaultImage))
	fmt.Println(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return domain.Response{}, err
	}
	defer func() {
		if err != nil {
			// コンテナを削除
			cmd = exec.Command("ssh", "-o", "StrictHostKeyChecking=no", "-i", sshKeyPath, vmUser+"@"+hostIP,
				fmt.Sprintf("docker rm -f %s", containerName))
			fmt.Println(cmd)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println("Failed to remove container:", err)
			}
		}
	}()

	guestIP := "10.10.10.10"

	// Dockerコンテナをネットワークに接続し、IPアドレスを付与
	cmd = exec.Command("ssh", "-o", "StrictHostKeyChecking=no", "-i", sshKeyPath, vmUser+"@"+hostIP,
		fmt.Sprintf("docker network connect --ip=%s mybridge %s", guestIP, containerName))
	fmt.Println(cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return domain.Response{}, err
	}

	// インスタンスの空きを更新（固定値ステータス）
	id, err := s.rep.CreateInstance(hostIP, containerName)
	if err != nil {
		return domain.Response{}, err
	}

	// レスポンスの返却
	response := domain.Response{
		ID: id,
		IP: guestIP,
	}

	fmt.Println("Instance created")
	return response, nil
}
