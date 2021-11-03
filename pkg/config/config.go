package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var GLOBAL_CONF = ServerConfig{}

type ServerConfig struct {
	Server struct {
		Port string `yaml:"port"`
	}
	DataBase struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Db       string `yaml:"db"`
		UserName string `yaml:"username"`
		PassWord string `yaml:"password"`
	}
	Redis struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
	WxSdk struct {
		Appid     string `yaml:"appid"`
		AppSecret string `yaml:"appsecret"`
	}
	TaoSdk struct {
		AppKey    string `yaml:"appkey"`
		AppSecret string `yaml:"appsecret"`
		Router    string `yaml:"router"`
		Adzoneid  string `yaml:"adzoneid"`
	}
}

func InitServerConfig() {
	GLOBAL_CONF = ServerConfig{}
	// 获取本地路径
	path, _ := os.Getwd()
	// 拼接读取配置文件
	yamlFile, err := ioutil.ReadFile(path + "/sys_config.yaml")
	if err != nil {
		log.Fatalln("ReadFile yaml error...")
		return
	}
	// 反射至结构体内
	err = yaml.Unmarshal(yamlFile, &GLOBAL_CONF)
	if err != nil {
		log.Fatalln("Unmarshal yaml error...")
		return
	}
}
