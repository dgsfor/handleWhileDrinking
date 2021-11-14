package conf

import (
	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 初始化配置
	SetupConfig()
	// 初始化redis
	SetupRedis()
}
