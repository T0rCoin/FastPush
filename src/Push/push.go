package Push

import (
	"BadOrange/conf"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

const(
	telegramAPI string = "https://api.telegram.org"
)

var(
	telegramBotAPI = telegramAPI + "/bot" + conf.Get("BotConfig","token") + "/"
	Client = &fasthttp.Client{
		MaxConnsPerHost: 10000,
		ReadTimeout: 400 * time.Millisecond,
		WriteTimeout: 400 * time.Millisecond,
	}
)

func SendMessage(Body string, ChatID int) []byte {

	url := fmt.Sprintf("%vsendmessage?text=%v&parse_mode=HTML&chat_id=%v", telegramBotAPI, Body, ChatID)

	_ , resp, err := Client.Get(nil, url)

	if err != nil {
	    fmt.Println(err,url)
		return []byte("")
	}

	return resp

}
