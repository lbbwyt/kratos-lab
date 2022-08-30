package config

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

/**
bootstrap.yaml
*/
var GConfig *BootstrapConfig

type NacosConfig struct {
	Ip    string `yaml:"ip"`
	Port  uint64 `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

type BootstrapConfig struct {
	Nacos NacosConfig `yaml:"nacos"`
}

func Boot(cfgPath string) {
	//默认读取本地配置
	op := file.NewSource(cfgPath)

	log.Info("bootstrap配置地址:" + cfgPath)

	//从配置中心读取配置
	//op = conf1.NewConfigSource(client, conf1.WithGroup("test"), conf1.WithDataID("config.yaml"))

	c := config.New(
		config.WithSource(
			op,
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc = &BootstrapConfig{}
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	GConfig = bc
}
