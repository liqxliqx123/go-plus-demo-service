package mysql

import (
	"gorm.io/gorm"
)

// mockgen -source=db/mysql/mysql.go -destination=db/mysql/mysql_mock.go -package=mysql

// xxxMySQLInterface is the interface of mysql
type xxxMySQLInterface interface {
}

// xxxMySQLClient is the struct of xxx client
type xxxMySQLClient struct {
	gormDB *gorm.DB
}

// NewxxxMySQLClient new xxx mysql db
func NewxxxMySQLClient(gormDB *gorm.DB) xxxMySQLInterface {
	return &xxxMySQLClient{
		gormDB: gormDB,
	}
}
