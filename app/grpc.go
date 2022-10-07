package app

import (
	"context"
	"my-demo-service/pb"
	"my-demo-service/server/grpc"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.xxx.com/xxx-xxx/go-kit/utils/grpcext"
	stdgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func initGRPCServerApplicationHook(o *Application) error {
	o.GRPCServer = grpcext.NewGRPCServer(o.Config.Service.Name)

	pb.RegisterServiceServer(o.GRPCServer, &grpc.ServiceServerImpl{})
	reflection.Register(o.GRPCServer)
	grpc_health_v1.RegisterHealthServer(o.GRPCServer, health.NewServer())

	gwmux := runtime.NewServeMux()

	addr := fmt.Sprintf("%s:%d", o.Config.Service.Host, o.Config.Service.Port)

	if err := pb.RegisterServiceHandlerFromEndpoint(context.Background(), gwmux, addr, []stdgrpc.DialOption{stdgrpc.WithInsecure()}); err != nil {
		return err
	}

	o.AddDestroyHook(func(o *Application) error {
		o.GRPCServer.GracefulStop()
		return nil
	})

	return nil
}
