package cmd

import (
	"my-demo-service/config"
	"my-demo-service/utils"

	kitConfig "gitlab.xxx.com/xxx-xxx/go-kit/utils/config"
)

// InitConfig loads config
func InitConfig() {
	//dir, _ := os.Getwd()
	//logger.Info(dir)
	//
	//p := apollo.Apollo{
	//	AppID:           viper.GetString("APOLLO.APP_ID"),
	//	Cluster:         viper.GetString("APOLLO.CLUSTER"),
	//	NameSpaceNames:  []string{viper.GetString("APOLLO.NAMESPACE")},
	//	MetaAddr:        viper.GetString("APOLLO.HOST"),
	//	AccessKeySecret: viper.GetString("APOLLO.TOKEN"),
	//	ServiceName:     viper.GetString("SERVICE.NAME"),
	//}
	//marshal, _ := json.Marshal(p)
	//logger.Info(string(marshal))

	kitConfig.LoadConf(utils.GetProjectPath(), config.GetConfig())
}
