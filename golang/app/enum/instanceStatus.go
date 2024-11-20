package enum

// 値の書き換えを防ぐため、実際に値を扱う構造体およびそのフィールドはprivateで宣言
type _instanceStatus struct {
	key   uint8
	value string
}

func (c _instanceStatus) Key() uint8 {
	return c.key
}

func (c _instanceStatus) Value() string {
	return c.value
}

type instanceStatus struct {
	NotStarted  _instanceStatus
	Running     _instanceStatus
	InOperation _instanceStatus
	Deleted     _instanceStatus
}

var InstanceStatus = instanceStatus{
	NotStarted:  _instanceStatus{key: 1, value: "停止中"},
	Running:     _instanceStatus{key: 2, value: "起動中"},
	InOperation: _instanceStatus{key: 3, value: "稼働中"},
	Deleted:     _instanceStatus{key: 4, value: "削除済み"},
}
