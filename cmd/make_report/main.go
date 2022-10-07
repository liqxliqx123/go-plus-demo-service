package main

import (
	"context"
	"my-demo-service/app"
	"my-demo-service/config"
	"my-demo-service/service/make_xxx"
	"os"
	"strconv"
	"time"

	"gitlab.xxx.com/xxx-xxx/go-kit/logger"
	kitConfig "gitlab.xxx.com/xxx-xxx/go-kit/utils/config"
)

func init() {
	kitConfig.LoadConf(".", config.GetConfig())
}

func main() {
	if len(os.Args) < 2 {

		return
	}
	atoi, err2 := strconv.Atoi(os.Args[1])
	if err2 != nil {
		logger.Errorf("wrong args, %v", os.Args[1])
		return
	}
	resultID := int64(atoi)
	jobEx := app.Default()
	defer jobEx.Shutdown()
	jobEx.DryRun()
	// do job

	xxx := make_xxx.NewInterfaceMakexxx(app.Default().GORMDB)
	cfg := config.GetConfig()
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.Service.MakexxxDataTimeout))
	defer cancelFunc()
	err := xxx.GetxxxData(ctx, resultID)
	if err != nil {
		logger.Errorf("%#v", err)
	}

}
