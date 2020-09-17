package db

import (
	"log"
	"testing"

	"github.com/robfig/config"
	"github.com/stretchr/testify/assert"
)

func TestConnectMysql(t *testing.T) {
	type args struct {
		host     string
		port     int32
		username string
		password string
		dbName   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "连接成功",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := config.ReadDefault("../jnvwa.cfg")
			if err != nil {
				log.Printf("err=%v", err)
				return
			}
			host, _ := c.String("mysql", "host")
			port, _ := c.Int("mysql", "port")
			username, _ := c.String("mysql", "username")
			password, _ := c.String("mysql", "password")
			dbName, _ := c.String("mysql", "dbname")
			mh := GetMysqlHandler(host, int32(port), username, password, dbName)
			db := mh.GetDB()
			assert.NotEqual(t, db, nil)
		})
	}
}
