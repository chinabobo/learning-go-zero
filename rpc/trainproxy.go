package main

import (
	"flag"
	"fmt"

	"github.com/chinabobo/learning-go-zero/rpc/internal/config"
	"github.com/chinabobo/learning-go-zero/rpc/internal/server"
	"github.com/chinabobo/learning-go-zero/rpc/internal/svc"
	"github.com/chinabobo/learning-go-zero/rpc/train_proxy"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/trainproxy.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		train_proxy.RegisterTrainProxySrvServer(grpcServer, server.NewTrainProxySrvServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
