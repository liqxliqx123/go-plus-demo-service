package beluga

import (
	"context"
	"my-demo-service/models"
)

// mockgen -source=vendor/gitlab.xxx.com/xxx-xxx/go-kit/services/beluga/client.go -destination=service/beluga/beluga_client_mock.go -package=beluga

var defaultBelugaService InterfaceBeluga

//InterfaceBeluga beluga service interface
type InterfaceBeluga interface {
	CreateTask(ctx context.Context, msg *models.xxxKafkaMsg, belugaTaskName string, tasks ...string) (int64, error)
}

//GetInterfaceBeluga get beluga service
func GetInterfaceBeluga() InterfaceBeluga {
	return defaultBelugaService
}
