package Push

import (
	"BadOrange/conf"
	"fmt"
	"github.com/valyala/fasthttp"
)

const(
	telegramAPI string = "https://api.telegram.org"
)

var(
	telegramBotAPI = telegramAPI + "/bot" + conf.Get("BotConfig","token") + "/"
)

func SendMessage(Body string, ChatID int) []byte {

	url := fmt.Sprintf("%vsendmessage?text=%v&parse_mode=HTML&chat_id=%v", telegramBotAPI, Body, ChatID)

	_ , resp, err := fasthttp.Get(nil, url)

	if err != nil {
		return []byte("")
	}

	return resp

}
