package conf

import (
	"BadOrange/src/logs"
	"fmt"
	"github.com/unknwon/goconfig"
	"os"
)

const(
	FilePath string = "./conf/conf.ini"
)

var RedisConf = map[string]string{
	"name":    "redis",
	"type":    "tcp",
	"address": Get("Redis","Address"),
	"auth":    Get("Redis","Password"),
}

func Get(section string, key string) string{

	cfg, err := goconfig.LoadConfigFile(FilePath)

	if err != nil{

		_ = logs.WriteLog("run.log", fmt.Sprintf("%s", err))

		fmt.Printf("Error:CANT LOAD CONFIG FILE,%v\n",err)

		os.Exit(1)

	}

	value, err := cfg.GetValue(section,key)

	if err != nil {

		_ = logs.WriteLog("run.log", fmt.Sprintf("%s", err))

		fmt.Printf("Error:Load Config Error%v\n",err)

		os.Exit(1)

	}

	return value

}