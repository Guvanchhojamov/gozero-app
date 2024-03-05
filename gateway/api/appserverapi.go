package main

import (
	"flag"
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/app"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/config"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/handler"
	"github.com/Guvanchhojamov/gozero-app/gateway/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/appserverapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	newApp, err := app.NewApp(c)
	if err != nil {
		logx.Error(err)
		return
	}
	ctx := svc.NewServiceContext(c, newApp)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
