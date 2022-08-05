package main

import (
	"fmt"
	"time"
)

func main() {
	Debug := false
	LogLevel := "Info"
	startUpTime := time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
