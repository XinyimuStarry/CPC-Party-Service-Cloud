package main

import "CPC_Party_Service_Cloud/rest"

func main() {
	// 日志：第一优先级保障
	rest.LOGGER.Init()
	defer rest.LOGGER.Destroy()
}
