package config

import (
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/appConf"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	App appConf.App
}
