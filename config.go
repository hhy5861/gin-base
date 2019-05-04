package main

import (
	"github.com/hhy5861/go-common/client"
	"github.com/hhy5861/go-common/common"
	"github.com/hhy5861/go-common/jwt"
	"github.com/hhy5861/go-common/logger"
	"github.com/hhy5861/go-common/mysql"
	"github.com/hhy5861/go-common/server"
)

type (
	Config struct {
		System   *server.Config      `yaml:"system"`
		Client   *client.Client      `yaml:"client"`
		Logger   *logger.Logger      `yaml:"logger"`
		Mysql    *mysql.MysqlConf    `yaml:"mysql"`
		Redis    *common.RedisConfig `yaml:"redis"`
		Security struct {
			Jwt *jwt.JwtConfig `yaml:"jwt"`
		} `yaml:"security"`
	}
)
