package db

import (
	"database/sql"
	"my-demo-service/config"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type logDatabaseSourceManager struct {
	logger *zap.Logger
	next   DatabaseSourceManager
}

// NewLogDatabaseSourceManagerMiddleware 生成日志数据库管理器中间件
func NewLogDatabaseSourceManagerMiddleware(log *zap.Logger) DatabaseSourceManagerMiddleware {
	return func(next DatabaseSourceManager) DatabaseSourceManager {
		return &logDatabaseSourceManager{
			logger: log,
			next:   next,
		}
	}
}

func (m *logDatabaseSourceManager) Open(name string, option *config.DatabaseOption) (err error) {
	defer func(b time.Time) {
		log := m.logger.With(
			zap.String("module", "DatabaseSourceManager"),
			zap.String("method", "Open"),
			zap.String("duration", time.Since(b).String()),
			zap.String("name", name),
			zap.Object("option", option),
		)

		if err != nil {
			log = log.With(zap.Error(err))
			log.Error("Open database error")
			return
		}
		log.Info("Open database success")
	}(time.Now())

	return m.next.Open(name, option)
}

func (m *logDatabaseSourceManager) GetMysqlDatabase(name string) (db *sql.DB, err error) {
	defer func(b time.Time) {
		log := m.logger.With(
			zap.String("module", "DatabaseSourceManager"),
			zap.String("method", "GetMysqlDatabase"),
			zap.String("duration", time.Since(b).String()),
			zap.String("name", name),
		)
		if err != nil {
			log = log.With(zap.Error(err))
			log.Error("GetMysqlDatabase error")
			return
		}
		log.Info("GetMysqlDatabase success")
	}(time.Now())
	return m.next.GetMysqlDatabase(name)
}

func (m *logDatabaseSourceManager) GetRedisDatabase(name string) (db *redis.Client, err error) {
	defer func(b time.Time) {
		log := m.logger.With(
			zap.String("module", "DatabaseSourceManager"),
			zap.String("method", "GetRedisDatabase"),
			zap.String("duration", time.Since(b).String()),
			zap.String("name", name),
		)
		if err != nil {
			log = log.With(zap.Error(err))
			log.Error("GetRedisDatabase error")
			return
		}
		log.Info("GetRedisDatabase success")
	}(time.Now())
	return m.next.GetRedisDatabase(name)
}

var _ DatabaseSourceManager = (*logDatabaseSourceManager)(nil)
