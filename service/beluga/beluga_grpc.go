package beluga

import (
	"context"
	"my-demo-service/c"
	"my-demo-service/config"
	"my-demo-service/models"
	"fmt"
	"strings"
	"time"

	belugapb "gitlab.xxx.com/xxx-xxx/go-kit/pb/beluga"
	pb "gitlab.xxx.com/xxx-xxx/go-kit/pb/beluga"
	"gitlab.xxx.com/xxx-xxx/go-kit/services/beluga"
	belugagrpc "gitlab.xxx.com/xxx-xxx/go-kit/services/beluga/grpc"
	"gitlab.xxx.com/xxx-xxx/go-kit/utils/grpcext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var maxExecutionTime = int64(5 * 60 * 60)

var errorStrategy = belugapb.ErrorStrategy{
	Deal:  belugapb.ErrorStrategy_STOP,
	TryOn: 3,
	// 30 min
	Interval: 60,
}

//GrpcClient grpc client
type GrpcClient struct {
	grpcClient      beluga.Client
	serviceName     string
	apolloAppID     string
	apolloHost      string
	apolloCluster   string
	apolloNameSpace string
	env             string
}

//InitBelugaGRPCClient inti beluga rpc client
func InitBelugaGRPCClient(cfg *config.Config) error {
	var opts []grpc.DialOption
	if cfg.EndPoints.BelugaService.GRPC.Insecure {
		opts = grpcext.NewInsecureDialOptions()
	} else {
		opts = grpcext.NewDefaultDialOptions()
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.EndPoints.BelugaService.GRPC.Host, cfg.EndPoints.BelugaService.GRPC.Port), opts...)
	if err != nil {
		return err
	}
	client, err := belugagrpc.NewBelugaClient(belugagrpc.Config{
		Env:     cfg.EndPoints.BelugaService.GRPC.Env,
		Timeout: time.Duration(cfg.Service.ConnectionTimeout) * time.Second,
		Client:  pb.NewServiceClient(conn),
	})
	if err != nil {
		return err
	}

	var env = config.GetConfig().Env
	if env != "production" {
		env = "stage01"
	}

	defaultBelugaService = &GrpcClient{
		grpcClient:      client,
		serviceName:     config.GetConfig().Apollo.APPID,
		apolloAppID:     config.GetConfig().Service.Name,
		apolloHost:      config.GetConfig().Apollo.Host,
		apolloCluster:   config.GetConfig().Apollo.Cluster,
		apolloNameSpace: config.GetConfig().Apollo.NameSpace,
		env:             env,
	}
	return nil
}

//CreateTask  create a xxx argo task
func (g *GrpcClient) CreateTask(ctx context.Context, data *models.xxxKafkaMsg, belugaTaskName string, tasks ...string) (int64, error) {
	cfg := config.GetConfig()
	cfg.CmdImageTag = ""
	steps, err := g.buildArgoTask(ctx, data, cfg.CmdImageTag, belugaTaskName, tasks...)
	if err != nil {
		return 0, err
	}
	var taskID int64
	var md map[string][]string
	switch g.env {
	case c.Production:
		md = map[string][]string{
			"authorization": {
				prodBelugaToken,
			},
		}
	default:
		md = map[string][]string{
			"authorization": {
				stageBelugaToken,
			},
		}
	}
	ctx = metadata.NewIncomingContext(ctx, md)
	var taskName = fmt.Sprintf("%s-%s", "request_id", strings.Join(tasks, "-"))
	taskID, err = g.grpcClient.CreateStatefulDagTask(ctx, taskName, steps, nil, nil, &errorStrategy, maxExecutionTime, &pb.ResourceStrategy{
		NodeLabels: map[string]string{
			"argo-xxx": "true",
		},
	})
	if err != nil {
		return 0, err
	}
	return taskID, nil
}
