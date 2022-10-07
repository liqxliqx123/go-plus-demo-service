package config

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
)

var globalConfig *Config
var configOnce sync.Once

// GetConfig 获取该服务配置
func GetConfig() *Config {
	configOnce.Do(func() {
		globalConfig = &Config{}
	})
	return globalConfig
}

// Config 该服务相关配置
type Config struct {
	Env                  string                    `mapstructure:"ENV"`
	Service              Service                   `mapstructure:"SERVICE"`
	LogLevel             string                    `mapstructure:"LOG_LEVEL"`
	LogFile              string                    `mapstructure:"LOG_FILE"`
	EnableApollo         bool                      `mapstructure:"ENABLE_APOLLO"`
	Databases            map[string]DatabaseOption `mapstructure:"DATABASES"`
	MINIO                MINIOOption               `mapstructure:"MINIO"`
	StorageAPIGatewayURL string                    `mapstructure:"STORAGE_API_GATEWAY_URL"`
	Kafka                KafkaOption               `mapstructure:"KAFKA"`
	EndPoints            Endpoints                 `mapstructure:"ENDPOINTS"`
	Apollo               Apollo                    `mapstructure:"APOLLO"`
	CmdImageTag          string                    `mapstructure:"CMD_IMAGE_TAG"`
}

// MINIOOption minio
type MINIOOption struct {
	Endpoint    string `mapstructure:"ENDPOINT"`
	MinioBucket string `mapstructure:"MINIO_BUCKET"`
	AccessID    string `mapstructure:"ACCESS_ID"`
	AccessKey   string `mapstructure:"ACCESS_KEY"`
	SSL         bool   `mapstructure:"SSL"`
}

// Service defines service configuration struct.
type Service struct {
	Name                  string `mapstructure:"NAME"`
	Host                  string `mapstructure:"HOST"`
	Port                  int    `mapstructure:"PORT"`
	ConnectionTimeout     int    `mapstructure:"CONNECTION_TIMEOUT"`
	MakexxxDataTimeout int    `mapstructure:"MAKE_xxx_DATA_TIMEOUT"`
	xxxConsumerTimeout int    `mapstructure:"xxx_CONSUMER_TIMEOUT"`
}

// DatabaseOption 数据库链接相关配置
type DatabaseOption struct {
	Driver       string        `mapstructure:"DRIVER"`
	Host         string        `mapstructure:"HOST"`
	Port         uint16        `mapstructure:"PORT"`
	Username     string        `mapstructure:"USERNAME"`
	Password     string        `mapstructure:"PASSWORD"`
	DBName       string        `mapstructure:"DBNAME"`
	Timezone     string        `mapstructure:"TIMEZONE"`
	Charset      string        `mapstructure:"CHARSET"`
	PoolSize     int           `mapstructure:"POOL_SIZE"`
	Timeout      time.Duration `mapstructure:"TIMEOUT"`
	ReadTimeout  time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT"`
}

// MarshalLogObject 用于将对象传递给日志 *zap.Logger
func (o DatabaseOption) MarshalLogObject(en zapcore.ObjectEncoder) error {
	en.AddString("driver", o.Driver)
	en.AddString("host", o.Host)
	en.AddUint16("port", o.Port)
	en.AddString("username", o.Username)
	en.AddString("password", "********")
	en.AddString("dbname", o.DBName)

	if len(o.Timezone) > 0 {
		en.AddString("timezone", o.Timezone)
	}

	if len(o.Charset) > 0 {
		en.AddString("charset", o.Charset)
	}

	if o.PoolSize > 0 {
		en.AddInt("pool_size", o.PoolSize)
	}

	if o.Timeout > 0 {
		en.AddString("timeout", o.Timeout.String())
	}

	if o.ReadTimeout > 0 {
		en.AddString("read_timeout", o.ReadTimeout.String())
	}

	if o.WriteTimeout > 0 {
		en.AddString("write_timeout", o.WriteTimeout.String())
	}

	return nil
}

// KafkaOption kafka producer configuration
type KafkaOption struct {
	Bootstrap         string `mapstructure:"BOOTSTRAP"`
	Retries           int    `mapstructure:"RETRIES"`
	xxxTopic       string `mapstructure:"xxx_TOPIC"`
	RegisterURL       string `mapstructure:"REGISTER_URL"`
	Version           string `mapstructure:"VERSION"`
	Assignor          string `mapstructure:"ASSIGNOR"`
	Oldest            bool   `mapstructure:"OLDEST"`
	Verbose           bool   `mapstructure:"VERBOSE"`
	ConsumerGroup     string `mapstructure:"CONSUMER_GROUP"`
	ChannelBufferSize int    `mapstructure:"CHANNEL_BUFFER_SIZE"`
}

// Endpoints endpoints
type Endpoints struct {
	//OrcaServerURL             string   `mapstructure:"ORCA_SERVER_URL"`
	SMNServiceURL             string   `mapstructure:"SMN_SERVICE_URL"`
	OrcaServer                Endpoint `mapstructure:"ORCA_SERVER"`
	SirenServer               Endpoint `mapstructure:"SIREN_SERVER"`
	BelugaService             Endpoint `mapstructure:"BELUGA_SERVICE"`
	DCTServiceURL             string   `mapstructure:"DCT_URL"`
	DCServiceURL              string   `mapstructure:"DC_SERVICE_URL"`
	WBIServerURL              string   `mapstructure:"WBI_SERVER_URL"`
	TransService              Endpoint `mapstructure:"TRANSLATE_SERVICE"`
	AutocatServerURL          string   `mapstructure:"AUTOCAT_SERVER_URL"`
	AutocatURLGID             string   `mapstructure:"AUTOCAT_URL_GID"`
	BacktrackAutocatServerURL string   `mapstructure:"BACKTRACK_AUTOCAT_SERVER_URL"`
	AMPServerURL              string   `mapstructure:"AMP_SERVER_URL"`
	xxxService                string   `mapstructure:"xxx_SERVER_URL"`
	AuthServer                Endpoint `mapstructure:"AUTH_SERVER"`
	RecommendSearchURL        string   `mapstructure:"RECOMMEND_SERVER_URL"`
	FEStaticURL               string   `mapstructure:"FE_STATIC_URL"`
}

// Endpoint defines endpoint struct.
type Endpoint struct {
	EnableGRPC bool         `mapstructure:"ENABLE_GRPC"`
	GRPC       GRPCEndpoint `mapstructure:"GRPC"`
}

// GRPCEndpoint defines grpc endpoint struct.
type GRPCEndpoint struct {
	Env      string `mapstructure:"ENV"`
	Insecure bool   `mapstructure:"INSECURE"`
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
}

// Apollo defines apollo struct
type Apollo struct {
	APPID       string `mapstructure:"APP_ID"`
	Cluster     string `mapstructure:"CLUSTER"`
	NameSpace   string `mapstructure:"NAMESPACE"`
	Host        string `mapstructure:"HOST"`
	ApolloToken string `mapstructure:"TOKEN"`
}

// ValidateConfig validates the config of gRPC
func (g *GRPCEndpoint) ValidateConfig() error {
	var err error = nil

	if len(g.Env) == 0 {
		err = fmt.Errorf("gRPC environment is not specified")
	}

	if len(g.Host) == 0 {
		err = fmt.Errorf("%w, gRPC HOST is not specified", err)
	}

	if len(g.Port) == 0 {
		err = fmt.Errorf("%w, gRPC PORT is not specified", err)
	}

	return err
}
