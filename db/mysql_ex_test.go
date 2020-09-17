package db

import (
	"testing"

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
		//{
		//name: "连接成功",
		//args: args{
		//host:     "10.129.50.89",
		//port:     3306,
		//username: "root",
		//password: "123456",
		//dbName:   "mybase",
		//},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mh := GetMysqlHandler(tt.args.host, tt.args.port, tt.args.username, tt.args.password, tt.args.dbName)
			db := mh.GetDB()
			assert.NotEqual(t, db, nil)
		})
	}
}
