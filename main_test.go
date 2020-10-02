package main

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	msg := `{"wxpush":{"appid":"xxx","secret":"xxx","touid":"xxx","tplid":"xxx"},"apps":[{"area":"cn","id":"1261944766","plat":"appstore"},{"area":"us","id":"932747118","plat":"appstore"},{"area":"us","id":"1137819437","plat":"appstore"}]}`
	event := TimerEvent{Message: msg}
	str, err := Exec(event)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("运行结果：%s\n", str)
}
