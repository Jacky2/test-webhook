package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func getConfig() {
	clientConfig := constant.ClientConfig{
		TimeoutMs:      10 * 1000, //http请求超时时间，单位毫秒
		NamespaceId:       "public", //nacos命名空间
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
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

	fmt.Println("GetConfig,config :" + content)

}