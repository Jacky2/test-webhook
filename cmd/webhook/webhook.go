package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"test-webhook/config"
	model "test-webhook/model"
	"test-webhook/notifier"
)

var (
	//h            bool
	//defaultRobot string

	vers *bool
	help *bool
	conf *string

	version = "No Version Provided"
)


func init() {
	//flag.BoolVar(&h, "h", false, "help")
	//flag.StringVar(&defaultRobot, "defaultRobot", "", "global dingtalk robot webhook, you can overwrite by alert rule with annotations dingtalkRobot")
	//flag.Parse()
	//
	//if h {
	//	flag.Usage()
	//	return
	//}

	vers = flag.Bool("v", false, "display the version.")
	help = flag.Bool("h", false, "print this help.")
	conf = flag.String("f", "", "specify configuration file.")

	flag.Parse()

	if *vers {
		fmt.Println("Version:", version)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

}


func main() {
	parseConf(*conf)
	netAddr := config.Config.Net.BindIp + ":" + strconv.Itoa(config.Config.Net.Port)

	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defaultRobotUrl := config.Config.Sender["robot"].API

		err = notifier.Send(notification, defaultRobotUrl)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})

	})
	router.Run(netAddr)
}


// auto detect configuration file
func parseConf(conf string) {
	if err := config.Parse(conf); err != nil {
		fmt.Println("cannot parse configuration file:", err)
		os.Exit(1)
	}
}