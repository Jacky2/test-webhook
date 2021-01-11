package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
	"strconv"
	"testing"
)

func parseConf() {
	conf := "/Volumes/Data/Code/github/test-webhook/etc/alert.yaml"
	if err := Parse(conf); err != nil {
		fmt.Println("cannot parse configuration file:", err)
		os.Exit(1)
	}
}

func TestGetconfig(t *testing.T) {
	parseConf()
	test := Config.Net.BindIp + ":" + strconv.Itoa(Config.Net.Port)
	fmt.Println(test)
	test2 := Config.Sender["robot"].API
	//test2 := Config.Config.Local.Enable
	fmt.Println(test2)
}

func TestGetNacosConfig(t *testing.T) {
	clientConfig := constant.ClientConfig{
		TimeoutMs:      10 * 1000, //http请求超时时间，单位毫秒
		NamespaceId:       "public", //nacos命名空间
		LogDir:              "~/.nacos/log",
		CacheDir:            "~/.nacos/cache",
		Username:            "nacos",
		Password:            "nacos", // the password for nacos auth
		UpdateThreadNum:   20, //更新服务的线程数
		NotLoadCacheAtStart: true, //在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true, //当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	// 至少一个
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "alert",
		Group:  "DEFAULT_GROUP",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("GetConfig,config :" + content)

	ParseString(content)
	test := Config.Net.BindIp + ":" + strconv.Itoa(Config.Net.Port)
	fmt.Println(test)
	test2 := Config.Sender["robot"].API
	//test2 := Config.Config.Local.Enable
	fmt.Println(test2)


}