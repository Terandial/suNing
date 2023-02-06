package dao

import (
	"github.com/jmoiron/sqlx"
)

// 定义一个全局对象db
var DB *sqlx.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"//本地数据库，就不写详细情况了哈，请见谅
	

	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil

}
