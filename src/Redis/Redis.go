package Xredis

import (
	. "BadOrange/conf"
	"BadOrange/src/logs"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var Client *redis.Pool

func Init()  {
	Client = &redis.Pool{
		MaxIdle: 16,
		MaxActive: 0,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(RedisConf["type"], RedisConf["address"])
			if err != nil{
				_ = logs.WriteLog("run.log", fmt.Sprintf("%s", err))
				return nil, err
			}
			if RedisConf["auth"] == "" {
				return c, nil
			}else {
				if _, err := c.Do("AUTH",RedisConf["auth"]); err != nil{
					c.Close()
					_ = logs.WriteLog("run.log", fmt.Sprintf("%s", err))
					return nil, err
				}
			}
			return c, err
		},
	}
}

func GetPushLastUIDObj() (interface{},error) {
	var err error
	var resp interface{}
	rds := Client.Get()
	defer rds.Close()
	_ , err = rds.Do("SELECT",13)
	resp, err = rds.Do("BRPOPLPUSH","push_queue", "push_queue_temp", 43600)
	if err != nil{
		_ = logs.WriteLog("run.log", fmt.Sprintf("%s", err))
		return nil, err
	}
	return resp, nil
}