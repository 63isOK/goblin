package main

import (
	"fmt"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)

func main() {
	c := &config.AppConfig{
		AppID:          "id",
		Cluster:        "dev",
		IP:             "ip",
		NamespaceName:  "global",
		IsBackupConfig: true,
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("初始化Apollo配置成功")

	// Use your apollo key to test
	cache := client.GetConfigCache(c.NamespaceName)
	// value, _ := cache.Get("zookeeper")
	// fmt.Println(value)

	cache.Range(func(key, value interface{}) bool {
		println(key, " - ", value)
		return true
	})
}
