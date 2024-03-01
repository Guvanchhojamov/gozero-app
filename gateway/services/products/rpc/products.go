package main

import (
	"flag"
	"fmt"

	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/internal/server"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/internal/svc"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/products/rpc/v1"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/products.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		v1.RegisterProductServiceServer(grpcServer, server.NewProductServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
