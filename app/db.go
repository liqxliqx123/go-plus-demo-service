package app

import (
	"my-demo-service/c"
	"my-demo-service/db"
	"fmt"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmsql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func initDatabasesApplicationHook(app *Application) error {
	loggerMiddleware := db.NewLogDatabaseSourceManagerMiddleware(app.Logger.Desugar())
	manager, err := db.NewDatabaseSourceManager(app.Config.Databases, loggerMiddleware)
	if err != nil {
		return err
	}
	app.DBManager = manager

	return nil
}

func initGormDatabasesApplicationHook(app *Application) error {
	mysqlConn, err := app.DBManager.GetMysqlDatabase(c.xxxMySQLDB)
	if err != nil {
		return err
	}

	conn, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlConn,
	}), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	_ = conn.Callback().Create().Before("gorm:create").Register("traceBeforeInsert", scopeBeforeFunc("INSERT"))
	_ = conn.Callback().Create().After("gorm:create").Register("traceAfterInsert", scopeAfterFunc())
	_ = conn.Callback().Delete().Before("gorm:delete").Register("traceBeforeDelete", scopeBeforeFunc("DELETE"))
	_ = conn.Callback().Delete().After("gorm:delete").Register("traceAfterDelete", scopeAfterFunc())
	_ = conn.Callback().Query().Before("gorm:query").Register("traceBeforeQuery", scopeBeforeFunc("SELECT"))
	_ = conn.Callback().Query().After("gorm:query").Register("traceAfterQuery", scopeAfterFunc())
	_ = conn.Callback().Update().Before("gorm:update").Register("traceBeforeUpdate", scopeBeforeFunc("UPDATE"))
	_ = conn.Callback().Update().After("gorm:update").Register("traceAfterUpdate", scopeAfterFunc())
	app.GORMDB = conn
	return nil
}

func scopeBeforeFunc(operation string) func(*gorm.DB) {
	return func(s *gorm.DB) {
		spanType := fmt.Sprintf("db.%s.%s", s.Name(), operation)

		span, _ := apm.StartSpan(s.Statement.Context, "", spanType)
		s.Set("apm-span", span)
	}
}

func scopeAfterFunc() func(*gorm.DB) {
	return func(s *gorm.DB) {
		value, ok := s.Get("apm-span")
		if ok {
			span := value.(*apm.Span)
			sql := s.Dialector.Explain(s.Statement.SQL.String(), s.Statement.Vars...)
			span.Name = apmsql.QuerySignature(sql)
			span.Context.SetDatabase(apm.DatabaseSpanContext{
				Instance:  s.Name(),
				Type:      "sql",
				Statement: sql,
			})
			span.Context.SetDatabaseRowsAffected(s.Statement.RowsAffected)
			span.End()
		}
	}
}
