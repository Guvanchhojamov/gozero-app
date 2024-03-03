package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/server"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/orders/rpc/v1"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/orders.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx, err := svc.NewServiceContext(c)
	if err != nil {
		logx.ErrorStack(err)
		return
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		v1.RegisterOrderServiceServer(grpcServer, server.NewOrderServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
