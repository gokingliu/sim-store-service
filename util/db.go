package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct{}

var db *sql.DB
var err error

// ConnDB 建立 DB 连接
func (s *DB) ConnDB() *sql.DB {
	if db == nil {
		dsn := "root:12345@tcp(127.0.0.1:3306)/sim" // 用户名:密码啊@tcp(ip:端口)/数据库名称
		db, err = sql.Open("mysql", dsn)            // 连接数据集，open 不会检验用户名和密码
		if err != nil {
			new(Log).LogError("配置 mysql 数据库错误", err)
		}
		err = db.Ping() // 尝试连接数据库
		if err != nil {
			new(Log).LogError("连接 mysql 错误", err)
		}
		db.SetMaxIdleConns(10) // 设置数据库连接池的最大连接数
	}

	return db
}
