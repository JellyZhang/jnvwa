package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	timeout = "10s" //连接超时，10秒
)

type mysqlHandler struct {
	host     string
	port     int32
	username string
	password string
	dbName   string
	timeout  string
	DB       *gorm.DB
}

// generate mysqlHandler
func GetMysqlHandler(host string, port int32, username, password, dbName string) *mysqlHandler {
	mh := &mysqlHandler{
		host:     host,
		port:     port,
		username: username,
		password: password,
		dbName:   dbName,
		timeout:  timeout,
	}
	mh.mustConnectMysql()
	return mh
}

// connect to Mysql
func (handler *mysqlHandler) connectMysql() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local&timeout=%s",
		handler.username,
		handler.password,
		handler.host,
		handler.port,
		handler.dbName,
		handler.timeout)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	handler.DB = db

	return nil
}

// must connect to Mysql, if fail, then panic
func (handler *mysqlHandler) mustConnectMysql() {
	err := handler.connectMysql()
	if err != nil {
		panic(err)
	}
}

// Get gorm.DB
func (handler *mysqlHandler) GetDB() *gorm.DB {
	return handler.DB
}
