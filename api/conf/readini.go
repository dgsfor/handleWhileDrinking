package conf

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"os"
	"strings"
)

var (
	EnvMode string
	Config *ini.File
)

func SetupConfig() {
	envParam := flag.String("env", "dev", "--env dev/prod")
	flag.Parse()

	EnvMode = *envParam

	if EnvMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	var err error
	Config, err = ini.Load("conf/config/app.ini", "conf/config/"+EnvMode+".ini")
	if err != nil {
		return
	}

	// 加入环境变量
	Config.ValueMapper = os.ExpandEnv

}
func GetConfig(key string) *ini.Key {
	parts := strings.Split(key, "::")
	section := parts[0]
	keyStr := parts[1]
	return Config.Section(section).Key(keyStr)

}

