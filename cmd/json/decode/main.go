//  //go:generate go run gen_config.go

package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const (
	aj = `
{
  "type":"aconfig",
  "config":{
    "name": "hi"
  }
}

`
	bj = `
{
  "type":"bconfig",
  "config":{
    "age":30
  }
}
`

	objj = `
{
  "type":"objconfig",
  "config":{
    "count": 12,
    "people": {
      "addr": "here",
      "age": 30
    }
  }
}
`
)

type ruleType struct {
	RType string `json:"type"`
}

type rule struct {
	Config any `json:"config"`
	ruleType
}

// 这种写法会导致无限循环: 如果结构体带泛型,那么解析时需要指定类型,这种场景显然不科学
//
// 最终的实现,还是从泛型+json走到了json+any, 在调用方处理不同的配置
func (r *rule) U(b []byte) error {
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	switch r.RType {
	case "aconfig":
		o := new(A)
		if err := mapstructure.Decode(r.Config, o); err != nil {
			return err
		}
		r.Config = o
	case "bconfig":
		o := new(B)
		if err := mapstructure.Decode(r.Config, o); err != nil {
			return err
		}
		r.Config = o
	case "objconfig":
		o := new(Obj)
		if err := mapstructure.Decode(r.Config, o); err != nil {
			return err
		}
		r.Config = o
	default:
		return errors.New("unknown type")
	}

	return nil

	// var a any
	// switch rt.RType {
	// case "aconfig":
	// 	a = new(rule[A])
	// case "bconfig":
	// 	a = new(rule[B])
	// case "objconfig":
	// 	a = new(rule[Obj])
	// default:
	// 	return errors.New("unknown type")
	// }

	// if err := json.Unmarshal(b, a); err != nil {
	// 	return err
	// }

	// *r = *a.(*rule[Config])
}

func parseJSON() {
	r := new(rule)
	if err := r.U([]byte(aj)); err != nil {
		panic(err)
	}
	fmt.Printf("type:%s, config: %+v\n", r.RType, r.Config)

	r = new(rule)
	if err := r.U([]byte(bj)); err != nil {
		panic(err)
	}
	fmt.Printf("type:%s, config: %+v\n", r.RType, r.Config)

	r = new(rule)
	if err := r.U([]byte(objj)); err != nil {
		panic(err)
	}

	fmt.Printf("type:%s, config: %+v\n", r.RType, r.Config)
}

type A struct {
	Name string `json:"name"`
}

type B struct {
	Age int `json:"age"`
}

type People struct {
	Addr string `json:"addr"`
	Age  int    `json:"age"`
}

type Obj struct {
	People People `json:"people"`
	Count  int    `json:"count"`
}

func main() {
	var o Obj
	if err := json.Unmarshal([]byte(objj), &o); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", o)

	fmt.Println("============")

	// parseRule([]byte(aj))
	// parseRule([]byte(bj))
	// parseRule([]byte(objj))
	fmt.Println("============")
	parseJSON()
}

// type Config interface {
// 	A | B | Obj
// }

// func parseRule(b []byte) (any, error) {
// 	var rt ruleType
// 	if err := json.Unmarshal(b, &rt); err != nil {
// 		return nil, err
// 	}

// 	var r any
// 	switch rt.RType {
// 	case "aconfig":
// 		r = new(rule[A])
// 	case "bconfig":
// 		r = new(rule[B])
// 	case "objconfig":
// 		r = new(rule[Obj])
// 	default:
// 		return nil, errors.New("unknown type")
// 	}

// 	err := json.Unmarshal(b, r)
// 	return r, err
// }
