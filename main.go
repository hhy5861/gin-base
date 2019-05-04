package main

import (
	"github.com/hhy5861/go-common/common"
	"github.com/hhy5861/go-common/logger"
	"github.com/hhy5861/go-common/mysql"
	"github.com/hhy5861/go-common/server"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"os"
)

type (
	Server struct {
		Config     *Config
		BaseServer *server.GinServer
	}
)

func NewServer(cfg *Config) *Server {
	return &Server{
		Config:     cfg,
		BaseServer: server.NewGinServer(cfg.System),
	}
}

func main() {
	svc := NewCommands().Commands()

	svc.CliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "run Commands",
			Action: func(c *cli.Context) {
				var cfg Config

				err := common.NewConfig(&cfg).ReadYaml(svc.PathFile)
				if err != nil {
					logger.Error(err)
				}

				db := mysql.NewMysql(cfg.Mysql).Connection()
				defer func() {
					err := db.Close()
					if err != nil {
						logger.Error(err)
					}
				}()

				svc := NewServer(&cfg)

				svc.BaseServer.I18nBundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
				svc.BaseServer.I18nBundle.MustLoadMessageFile("config/locales/en-US.yaml")
				svc.BaseServer.I18nBundle.MustLoadMessageFile("config/locales/zh-CN.yaml")

				svc.BaseServer.LoopCall(NewManufacture(&cfg))
				svc.BaseServer.RegisterRoute = svc.Router
				svc.BaseServer.GrpcServer.RegisteGrpcServer = svc.RegisteGrpcRouter

				svc.BaseServer.Run()
			},
		},
	}

	if err := svc.CliApp.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
