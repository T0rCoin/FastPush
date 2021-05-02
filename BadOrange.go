package main

import (
	router "BadOrange/app/route"
	"BadOrange/conf"
	Xredis "BadOrange/src/Redis"
	"BadOrange/src/logs"
	"fmt"
	"os"
)

func main()  {

	fmt.Printf("\n %c[1;40;33m%s%c[0m\n\n", 0x1B, "Welcome Use U.S. FireEye Inc. Telegram Fast Push Bot", 0x1B)

	route := router.SetUpRouter()

	Xredis.Init()

	if err := route.Run(conf.Get("Server","BindAddress")); err != nil {
		fmt.Println(err)
		_ = logs.WriteLog("run.log", fmt.Sprintf("%s", err))
		os.Exit(0)
	}

}