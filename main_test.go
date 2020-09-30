package main

import (
	"fmt"
	"log"
	"testing"
)

func Test_unmarshal(t *testing.T) {
	apps, err := unmarshal(`{"as":["123","1234"],"ps":["hell","world"]}`)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("应用：", apps)
}

func TestExec(t *testing.T) {
	msg := `{"wxpush":{"appid":"xxx","secret":"xxx","touid":"xxx","tplid":"xxx"},"apps":{"1261944766":"cn","932747118":"us"}}`
	event := TimerEvent{Message: msg}
	str, err := Exec(event)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("运行结果：%s\n", str)
}
