package db

import (
	"database/sql"
	"my-demo-service/config"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
)

// NewDatabaseSourceManager 从数据库配置中打开相关的数据库连接
func NewDatabaseSourceManager(options map[string]config.DatabaseOption, middlewares ...DatabaseSourceManagerMiddleware) (DatabaseSourceManager, error) {
	var m DatabaseSourceManager
	m = &basicDatabaseSourceManager{
		mysqlDatabases: make(map[string]*sql.DB),
		redisDatabases: make(map[string]*redis.Client),
	}

	for _, middleware := range middlewares {
		m = middleware(m)
	}

	for name, option := range options {
		if err := m.Open(name, &option); err != nil {
			return nil, err
		}
	}

	return m, nil
}

type basicDatabaseSourceManager struct {
	mu             sync.RWMutex
	mysqlDatabases map[string]*sql.DB
	redisDatabases map[string]*redis.Client
}

func (m *basicDatabaseSourceManager) Open(name string, option *config.DatabaseOption) (err error) {
	name = strings.ToLower(name)

	switch option.Driver {
	case DriverMysql:
		err = m.openMysqlDatabase(name, option)
		break
	case DriverRedis:
		err = m.openRedisDatabase(name, option)
		break
	default:
		err = fmt.Errorf("DatabaseSourceManager: unknown driver [%s]", option.Driver)
	}

	return
}

func (m *basicDatabaseSourceManager) openMysqlDatabase(name string, option *config.DatabaseOption) (err error) {
	name = strings.ToLower(name)

	m.mu.RLock()
	if _, has := m.mysqlDatabases[name]; has {
		m.mu.RUnlock()
		err = fmt.Errorf("DatabaseSourceManager: mysql database [%s] already exists", name)
		return
	}
	m.mu.RUnlock()

	var loc = time.Local

	if len(option.Timezone) > 0 {
		if loc, err = time.LoadLocation(option.Timezone); err != nil {
			return
		}
	}

	c := mysql.Config{
		User:                 option.Username,
		Passwd:               option.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", option.Host, option.Port),
		DBName:               option.DBName,
		Loc:                  loc,
		Timeout:              option.Timeout,
		ReadTimeout:          option.ReadTimeout,
		WriteTimeout:         option.WriteTimeout,
		ParseTime:            true,
		CheckConnLiveness:    true,
		AllowNativePasswords: true,
		MaxAllowedPacket:     4 << 20, // 4MB
		Collation:            "utf8mb4_general_ci",
	}

	if len(option.Charset) > 0 {
		c.Params = make(map[string]string)
		c.Params["charset"] = option.Charset
	}

	var db *sql.DB
	if db, err = sql.Open(DriverMysql, c.FormatDSN()); err != nil {
		return
	}

	// Set connection pool
	if option.PoolSize > 0 {
		db.SetMaxIdleConns(option.PoolSize)
		db.SetMaxOpenConns(option.PoolSize)
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	m.mysqlDatabases[name] = db

	return
}

func (m *basicDatabaseSourceManager) openRedisDatabase(name string, option *config.DatabaseOption) (err error) {
	name = strings.ToLower(name)

	m.mu.RLock()
	if _, has := m.redisDatabases[name]; has {
		m.mu.RUnlock()
		err = fmt.Errorf("DatabaseSourceManager: redis database [%s] already exists", name)
		return
	}
	m.mu.RUnlock()

	opt := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", option.Host, option.Port),
		Username:     option.Username,
		Password:     option.Password,
		ReadTimeout:  option.ReadTimeout,
		WriteTimeout: option.WriteTimeout,
		PoolSize:     option.PoolSize,
	}

	if len(option.DBName) > 0 {
		var dbName int
		if dbName, err = strconv.Atoi(option.DBName); err != nil {
			return
		}
		opt.DB = dbName
	}

	db := redis.NewClient(opt)

	m.mu.Lock()
	defer m.mu.Unlock()
	m.redisDatabases[name] = db
	return
}

func (m *basicDatabaseSourceManager) GetMysqlDatabase(name string) (*sql.DB, error) {
	name = strings.ToLower(name)

	if m.mysqlDatabases == nil {
		m.mu.Lock()
		m.mysqlDatabases = make(map[string]*sql.DB)
		m.mu.Unlock()
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	db, has := m.mysqlDatabases[name]
	if !has {
		return nil, fmt.Errorf("DatabaseSourceManager: mysql database [%s] doesn't exists", name)
	}

	return db, nil
}

func (m *basicDatabaseSourceManager) GetRedisDatabase(name string) (*redis.Client, error) {
	name = strings.ToLower(name)

	if m.redisDatabases == nil {
		m.mu.Lock()
		m.redisDatabases = make(map[string]*redis.Client)
		m.mu.Unlock()
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	db, has := m.redisDatabases[name]
	if !has {
		return nil, fmt.Errorf("DatabaseSourceManager: redis database [%s] doesn't exists", name)
	}

	return db, nil
}

// Type check
var _ DatabaseSourceManager = (*basicDatabaseSourceManager)(nil)
