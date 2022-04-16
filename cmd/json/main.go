package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// one := oneIntanceConfig{
	//   id:     1,
	//   config: 2,
	// }

	i := make([]int, 0)
	i = append(i, 1)
	i = append(i, 2)

	configs := instanceConfigs{
		// instances: i,
		DD: 2,
	}

	b, err := json.Marshal(configs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}

type oneIntanceConfig struct {
	id     int `json:"id"`
	config int `json:"config"`
}

// instanceConfigs 实例配置列表的定义
type instanceConfigs struct {
	// instances []int `json:"instances"`
	DD int `json:"dd"`
}
