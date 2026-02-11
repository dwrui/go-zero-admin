package main

import (
	"flag"
	"fmt"

	"system/internal/config"
	accountserviceServer "system/internal/server/accountservice"
	deptserviceServer "system/internal/server/deptservice"
	logserviceServer "system/internal/server/logservice"
	roleserviceServer "system/internal/server/roleservice"
	ruleserviceServer "system/internal/server/ruleservice"
	"system/internal/svc"
	"system/system"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/system.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		system.RegisterRuleServiceServer(grpcServer, ruleserviceServer.NewRuleServiceServer(ctx))
		system.RegisterDeptServiceServer(grpcServer, deptserviceServer.NewDeptServiceServer(ctx))
		system.RegisterRoleServiceServer(grpcServer, roleserviceServer.NewRoleServiceServer(ctx))
		system.RegisterAccountServiceServer(grpcServer, accountserviceServer.NewAccountServiceServer(ctx))
		system.RegisterLogServiceServer(grpcServer, logserviceServer.NewLogServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
