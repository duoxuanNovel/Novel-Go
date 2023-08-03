package model

import "time"

type IntTime int

const (
	LocalDateTimeFormat string = "2006-01-02T15:04:05.999Z"
)

// 将IntTime时间戳转换成字符串
func (t *IntTime) String() string {
	return time.Unix(int64(*t), 0).Format(LocalDateTimeFormat)
}

func (t *IntTime) MarshalJSON() ([]byte, error) {
	if *t == 0 {
		return []byte("\"\""), nil
	}
	return []byte("\"" + t.String() + "\""), nil
}
