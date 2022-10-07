package grpc

import (
	"context"

	"my-demo-service/pb"
)

// Status status api
func (s *ServiceServerImpl) Status(ctx context.Context, req *pb.Empty) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{Status: "ok"}, nil
}
