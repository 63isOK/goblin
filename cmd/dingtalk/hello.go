package main

import (
	"net/http"
	"strings"
)

func main() {
	msg := "机器人:测试消息"

	sendDingMsg(msg)
}

func sendDingMsg(msg string) {
	// 请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=a576d7f87b93bb2cde1a345210e146cc03fdc3857d87ae808819d1c5ac81b75d`
	content := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
	// 创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		// handle error
	}

	client := &http.Client{}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// 发送请求
	resp, err := client.Do(req)
	// 关闭请求
	defer resp.Body.Close()

	if err != nil {
		// handle error
	}
}
