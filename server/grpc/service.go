package grpc

import (
	"my-demo-service/pb"
)

// ServiceServerImpl implements grpc service interface
type ServiceServerImpl struct {
	pb.UnimplementedServiceServer
}
