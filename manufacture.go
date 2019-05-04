package main

import (
	"fmt"
	"github.com/hhy5861/go-common/client"
	"github.com/hhy5861/go-common/common"
	"github.com/hhy5861/go-common/google_grpc"
	"github.com/hhy5861/go-common/logger"
)

type (
	Manufacture struct {
		config *Config
	}
)

func NewManufacture(config *Config) *Manufacture {
	return &Manufacture{
		config: config,
	}
}

func (m *Manufacture) NewLogger() {
	logger.NewLogger(m.config.Logger)
}

func (m *Manufacture) NewRedis() {
	common.NewRedisStore(m.config.Redis, common.NewSerializeJson())
}

func (m *Manufacture) NewClient() {
	client.NewClient(m.config.Client)
}

func (m *Manufacture) NewGrpcClient() {
	google_grpc.NewGrpcClientService(&google_grpc.GrpcConfig{
		Cfg: map[string][]string{
			"default": {
				fmt.Sprintf("%s:%d", m.config.System.Host, m.config.System.Port),
			},
		},
	}).GrpcDial()
}
