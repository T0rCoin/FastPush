package fuse

import (
	"BadOrange/src/Push"
	Xredis "BadOrange/src/Redis"
	"BadOrange/src/logs"
	"fmt"
	"strconv"
	"time"
)

var Start = make(chan bool)

func PushAllStart(text string) bool {

	for i := 0; i < 1000; i++ {//起1000个goroutine

		time.Sleep(time.Microsecond * 100)
        
		go ThreadPush(text)

	}
    fmt.Println("Push Thread Read -> Start")
	close(Start)

	_ = logs.WriteLog("run.log","Push Start")

	return true

}

func ThreadPush(text string)  {//可以起一个独立的 goroutine 去推送
	<-Start
	for {
		obj, err := Xredis.GetPushLastUIDObj()
		if err != nil {
			fmt.Println(err)
			_ = logs.WriteLog("run.log", fmt.Sprintf("%s", err))
		}
		if obj == nil {
			break
		}
		ChatID, _ := strconv.Atoi(fmt.Sprintf("%s",obj))
		Push.SendMessage(text, ChatID)
	}
	return
}
