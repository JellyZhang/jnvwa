package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	timeout = "10s" //连接超时，10秒
)

// 连接Mysql
func ConnectMysql(host string, port int32, username, password, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local&timeout=%s",
		username,
		password,
		host,
		port,
		dbName,
		timeout)

	log.Printf("[jnvwa] connect mysql with host=%s, port=%d, username=%s, password=%s, dbName=%s", host, port, username, password, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
