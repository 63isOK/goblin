package main

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

func main() {
	showResponse()

	println("================")

	scanResponse()
}

func scanResponse() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func showResponse() {
	c := `
curl --location --request GET 'http://127.0.0.1:5320/parking_space_status_change' \
--header 'Content-Type: application/json' \
--data-raw '{
	"data": {
		"createTime": "2021-09-13 16:24:01",
		"electric": "3.5",
		"isOnline": 0,
		"id": "A10208CA",
		"state": 1,
		"heartReTime": "2021-09-13 16:23:53",
		"tag": 1,
		"uuid": "uuid123",
		"calibration": 0,
		"parkReTime": "2021-09-13 16:23:53"
	},
	"client": "PC",
	"version": "1.0.0",
	"user": "ChangHong",
	"key": 1631521470004
}'
`

	c = strings.ReplaceAll(c, "uuid123", uuid.New().String())
	cmd := exec.Command("sh", "-c", c)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	println(string(stdoutStderr))
}
