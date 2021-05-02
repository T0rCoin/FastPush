package service

import (
	"BadOrange/conf"
	"BadOrange/src/IP"
	PushX "BadOrange/src/Push"
	"BadOrange/src/fuse"
	"BadOrange/src/logs"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PushSignal struct {
	Token     string `json:"token"`
	Timestamp int    `json:"Timestamp"`
	Text      string `json:"text"`
}

type aPush struct {
	Token     string `json:"token"`
	Timestamp int    `json:"Timestamp"`
	UID       int `json:"Uid"`
	Text      string `json:"text"`
}

func SHA256(str string) string {
	return fmt.Sprintf("%x",sha256.Sum256([]byte(str)))
}

func Root(c * gin.Context)  {

	c.JSON(201, gin.H{
		"code": 201,
		"message": "Buffalo distributed telegram rover message push node",
	})

}

func PushAll(c * gin.Context)  {

	var Signal PushSignal

	if err := c.ShouldBindJSON(&Signal); err != nil {
		_ = logs.WriteLog("http.log", fmt.Sprintf("Bad Request-IP:%s-Message:%s", IP.GetRealIp(c.Request), err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"msg": err.Error(),
		})
		return
	}

	if Signal.Timestamp < int(time.Now().Unix() - 1) {
		_ = logs.WriteLog("http.log", fmt.Sprintf("Bad Request-IP:%s-Message:TimeStamp Error", IP.GetRealIp(c.Request)))
		c.JSON(400, gin.H{
			"status": 400,
			"msg": "Time is wrong, time zone is out of sync",
		})
		return
	}

	if SHA256(Signal.Token) != SHA256(conf.Get("Safe","PushToken")) {
		_ = logs.WriteLog("http.log", fmt.Sprintf("Bad Request-IP:%s-Message:Authentication failed", IP.GetRealIp(c.Request)))
		c.JSON(401, gin.H{
			"status": 401,
			"msg": "Could not verify your identity",
		})
		return
	}

	if status := fuse.PushAllStart(Signal.Text); !status {
		_ = logs.WriteLog("http.log", fmt.Sprintf("Bad Request-IP:%s-Message:Can't Push", IP.GetRealIp(c.Request)))
		c.JSON(503, gin.H{
			"status": 503,
			"msg": "Can't Push",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"msg": "success",
	})
	_ = logs.WriteLog("http.log", fmt.Sprintf("Success Request-IP:%s", IP.GetRealIp(c.Request)))
	return
}

func Push(c * gin.Context)  {

	var Signal aPush

	if err := c.ShouldBindJSON(&Signal); err != nil {
		_ = logs.WriteLog("http.log", fmt.Sprintf("Bad Request-IP:%s-Message:%s", IP.GetRealIp(c.Request), err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"msg": err.Error(),
		})
		return
	}

	if Signal.Timestamp < int(time.Now().Unix() - 1) {
		_ = logs.WriteLog("http.log", fmt.Sprintf("Bad Request-IP:%s-Message:TimeStamp Error", IP.GetRealIp(c.Request)))
		c.JSON(400, gin.H{
			"status": 400,
			"msg": "Time is wrong, time zone is out of sync",
		})
		return
	}

	if SHA256(Signal.Token) != SHA256(conf.Get("Safe","PushToken")) {
		_ = logs.WriteLog("http.log", fmt.Sprintf("Bad Request-IP:%s-Message:Authentication failed", IP.GetRealIp(c.Request)))
		c.JSON(401, gin.H{
			"status": 401,
			"msg": "Could not verify your identity",
		})
		return
	}

	PushX.SendMessage(Signal.Text, Signal.UID)
	c.JSON(200, gin.H{
		"status": 200,
		"msg": "success",
	})
	_ = logs.WriteLog("http.log", fmt.Sprintf("Success Request-IP:%s", IP.GetRealIp(c.Request)))

}