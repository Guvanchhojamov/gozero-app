package config

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/appConfig"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	App appConfig.App
}
