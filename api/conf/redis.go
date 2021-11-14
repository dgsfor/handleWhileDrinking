package conf

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

var RedisClient *redigo.Pool

func SetupRedis() {
	redis_host := GetConfig("redis::host")
	redis_port := GetConfig("redis::port")
	redis_auth := GetConfig("redis::password")
	// 建立连接池
	timeout, _ := GetConfig("redis::timeout").Duration()
	RedisClient = &redigo.Pool{
		MaxIdle:     16,                    //最初的连接数量
		MaxActive:   0,                     //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: timeout * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redigo.Conn, error) { //要连接的redis数据库
			c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%s", redis_host, redis_port))
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", redis_auth); err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}
