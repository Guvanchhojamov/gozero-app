package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Services Services
	App      App
}
type Services struct {
	Authorization zrpc.RpcClientConf
	Product       zrpc.RpcClientConf
	Order         zrpc.RpcClientConf
}

type App struct {
	DateBase     DateBase
	JWTSecretKey string
	Cache        Cache
}

type DateBase struct {
	Postgres Postgres
}
type Postgres struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	SslMode  string
}

type Cache struct {
	Host string
	Pass string
}
