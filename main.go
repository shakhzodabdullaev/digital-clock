package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gosuri/uilive"
)

func addZero(target *int) string {
	var tempTime string
	if *target < 10 {
		tempTime = "0" + strconv.Itoa(*target)
		return tempTime
	}
	tempTime = strconv.Itoa(*target)
	return tempTime
}

func main() {
	startUI := uilive.New()
	startUI.Start()

	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()
	currentSecond := currentTime.Second()

	fmt.Println("Clock has stopped")
	startUI.Stop()
}
