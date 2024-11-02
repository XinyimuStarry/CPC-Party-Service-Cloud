package utils

import "time"

// 一天中的第一秒钟
func FirstSecondOfDay(day time.Time) time.Time {
	local, _ := time.LoadLocation("Local")
	return time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, local)
}

// 明天此刻的时间
func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

// 昨天此刻的时间
func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// 将一个 时间 字符串转换成 时间 对象(yyyy-MM-dd HH:mm:ss)
func ConvertTimeToTimeString(time time.Time) string {
	return time.Local().Format("15:04:05")
}

// 将一个 日期+时间 字符串转换成 日期+时间 对象(yyyy-MM-dd HH:mm:ss)
func ConvertTimeToDateTimeString(time time.Time) string {
	return time.Local().Format("2006-01-02 15:04:05")
}
