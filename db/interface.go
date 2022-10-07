package db

import (
	"database/sql"
	"my-demo-service/config"

	"github.com/go-redis/redis/v8"
)

// DatabaseSourceManager 数据库管理器
type DatabaseSourceManager interface {
	Open(name string, option *config.DatabaseOption) error
	GetMysqlDatabase(name string) (*sql.DB, error)
	GetRedisDatabase(name string) (*redis.Client, error)
}

// DatabaseSourceManagerMiddleware 数据库管理器中间件
type DatabaseSourceManagerMiddleware func(DatabaseSourceManager) DatabaseSourceManager
