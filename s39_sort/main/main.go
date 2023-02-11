package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("%v %v %v\n", now.Year(), now.Day(), now.UnixMilli())

	add := now.Add(time.Hour * 12)
	fmt.Printf("%v %v %v\n", add.Year(), add.Day(), add.UnixMilli())

	fmt.Printf(now.Format("2006/01/02 15:04"))

	loc, _ := time.LoadLocation("Asia/Shanghai")
	time, _ := time.ParseInLocation("2006/01/02 15:04:05", "2022/01/02 15:04:05", loc)
	fmt.Printf("%v\n", time)
}
