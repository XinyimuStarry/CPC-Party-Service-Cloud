package utils

import "CPC_Party_Service_Cloud/rest"

// 带有panic恢复的方法

func PanicHandler() {
	if err := recover(); err != nil {
		rest.LOGGER.Error("异步任务错误: %v", err)
	}
}
func SafeMethod(f func()) {
	defer PanicHandler()
	//执行函数
	f()
}
