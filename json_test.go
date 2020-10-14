package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

// 定义json字段的struct
type User1 struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}

type User2 struct {
	Name     string    `json:"name"`
	Age      int       `json:"age,omitempty"` //忽略默认值
	Birthday time.Time `json:"birthday"`
}

// 序列化
func TestMarshal(t *testing.T) {
	u := User1{
		Name:     "cm",
		Age:      18,
		Birthday: time.Now(),
	}
	bs, _ := json.Marshal(u)
	js := string(bs)
	fmt.Println(js)
}

// 反序列化
func TestUnMarshal(t *testing.T) {
	js := "{\"Name\":\"cm\",\"Age\":18,\"Birthday\":\"2020-10-14T20:24:50.84444+08:00\"}"
	bs := []byte(js)
	u := User{}
	json.Unmarshal(bs, &u)
	fmt.Println(u)
}

// 序列化
func TestOmitempty(t *testing.T) {
	u := User2{
		Name:     "cm",
		Birthday: time.Now(),
	}
	bs, _ := json.Marshal(u)
	js := string(bs)
	fmt.Println(js)
}
