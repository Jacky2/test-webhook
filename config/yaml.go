package config

import (
	"fmt"
	"test-webhook/utils/file"
)

type ConfigT struct {
	GetConfig getConfgSection `yaml:"getConfig"`
	Net netSection `yaml:"net"`
	Sender map[string]senderSection `yaml:"sender"`
	WeChat wechatSection            `yaml:"wechat"`
	Email  emailSection             `yaml:"email"`
}

type netSection struct {
	Port int `yaml:"port"`
	BindIp string `yaml:"bindIp"`
}

type getConfgSection struct {
	Local localSection `yaml:"local"`
	Nacos nacosSection `yaml:"nacos"`
}

type localSection struct {
	Enable bool `yaml:"enable"`
}

type nacosSection struct {
	Enable bool `yaml:"enable"`
	IpAddr string `yaml:"ipAddr"`
	Port int `yaml:"port"`
	ContextPath string `yaml:"contextPath"`
}

type emailSection struct {
	SMTPHost  string `yaml:"smtpHost"`
	SMTPPort  int `yaml:"smtpPort"`
	SMTPUser  string `yaml:"smtpUser"`
	SMTPPass  string `yaml:"smtpPass"`
	SMTPInsecureSkipVerify  bool `yaml:"smtpInsecureSkipVerify"`
}

type wechatSection struct {
	CorpID  string `yaml:"corpId"`
	AgentID int    `yaml:"agentId"`
	Secret  string `yaml:"secret"`
}

type senderSection struct {
	Way    string `yaml:"way"`
	Worker int    `yaml:"worker"`
	API    string `yaml:"api"`
}

var Config *ConfigT

// Parse configuration file
func Parse(conf string) error {
	ymlFile := getYmlFile(conf)
	if ymlFile == "" {
		return fmt.Errorf("configuration file not found")
	}

	var c ConfigT
	err := file.ReadYaml(ymlFile, &c)
	if err != nil {
		return fmt.Errorf("cannot read yml[%s]: %v", ymlFile, err)
	}

	Config = &c
	return nil
}

func getYmlFile(conf string) string {
	yml := conf
	if file.IsExist(yml) {
		return yml
	}

	yml = "etc/alert.local.yml"
	if file.IsExist(yml) {
		return yml
	}

	yml = "etc/alert.yml"
	if file.IsExist(yml) {
		return yml
	}

	return ""
}

func ParseString(data string) error {
	var c ConfigT
	err := file.ReadYamlString(data, &c)
	if err != nil {
		return fmt.Errorf("cannot read yml %v", err)
	}

	Config = &c
	return nil
}
