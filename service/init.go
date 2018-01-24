package service

import (
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

//注册模型
func init() {
	initDb()
}



