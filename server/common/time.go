package common

import "time"

// NowTime 时间格式化
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

