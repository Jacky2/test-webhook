package main

import (
	"fmt"
	"test-webhook/config"
	"testing"
)


func TestGetConfig(t *testing.T) {

	parseConf("/Volumes/Data/Code/github/test-webhook/etc/alert.yaml")


	fmt.Println(config.Config.Sender["im"].API)
	//test := config.Config.Email.SMTPHost
	//fmt.Println(test)
	//test2 := config.Config.Sender["robot"].API
	//fmt.Println(test2)
}
