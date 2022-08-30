package main

import (
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"helloworld/internal/conf"
	bootstrap "helloworld/pkg/config"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "helloworld"
	// Version is the version of the compiled software.
	Version = "v0.0.1"
	// flagconf is the config flag.
	flagconf string

	//bootstrap config
	bootstrapConf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&bootstrapConf, "bootstrap", "../../pkg/config/bootstrap.yaml", "config path, eg: -bootstrap bootstrap.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, registry *nacos.Registry) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(registry),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	//读取bootstrap 配置
	bootstrap.Boot(bootstrapConf)

	var op config.Source
	if bootstrap.GConfig.Nacos.Debug {
		//测试环境读取本地配置
		op = file.NewSource(flagconf)
	} else {
		//从配置中心加载配置
		log.Infof("nacos bootstrap config : %v", bootstrap.GConfig)

		op = wireConfigSource(bootstrap.GConfig.Nacos.Ip, bootstrap.GConfig.Nacos.Port, bootstrap.WithGroup("prod"), bootstrap.WithDataID("config.yaml"))
	}

	c := config.New(
		config.WithSource(
			op,
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	log.Infof("app bootstrap config : %v", bc)
	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Registry.GetNacos(), logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
