package domain

type RequestParam struct {
	SshKeyId uint32 `json:"ssh_key_id"`
}

type HostAvailability struct {
	HostIP    string `gorm:"column:host_ip"`
	Available bool   `gorm:"column:is_available"`
	Count     int    `gorm:"column:count"`
}

type Response struct {
	ID uint32 `json:"id"`
	IP string `json:"ip"`
}

// 該当データなしエラー格納用
type NotFoundError struct {
	error
}
