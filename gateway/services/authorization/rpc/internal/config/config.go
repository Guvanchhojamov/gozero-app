package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	App App
}
type App struct {
	Postgres     Postgres
	JWTSecretKey string
	PasswordSalt string
	Cache        Cache
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
