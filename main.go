package main

import (
	"gojekyll/cmd"
	"log"
	"time"
)

var buildTime, gitHash string

func init() {
	l, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("chinese time zone")
		time.Local = l
	}
}
func main() {
	cmd.Execute(buildTime, gitHash)
}
