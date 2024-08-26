package timeutils

import "time"

// GetNowTime 获取当前格式化时间
func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 计算时间差,秒数

func GetTimeDiff(startTime, endTime string) int {
	st, _ := time.Parse("2006-01-02 15:04:05", startTime)
	et, _ := time.Parse("2006-01-02 15:04:05", endTime)
	return int(et.Sub(st).Seconds())
}
