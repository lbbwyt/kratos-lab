package registry

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"helloworld/internal/conf"
	"log"
)

func NewNacosRegistry(cfg *conf.Registry_Nacos) *nacos.Registry {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(cfg.GetIp(), cfg.GetPort()),
	}

	//cc := constant.ClientConfig{
	//	NamespaceId:         "public",
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "/tmp/nacos/log",
	//	CacheDir:            "/tmp/nacos/cache",
	//	LogLevel:            "debug",
	//}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			//ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Panic(err)
	}
	return nacos.New(client)
}
