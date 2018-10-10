package main

import (
	"coding.net/askuy/oauth/backend/app/model"
	"github.com/gin-gonic/gin"
	"github.com/godefault/caller"
	"github.com/godefault/caller/ginsession"
	"github.com/godefault/caller/gorm"
	"github.com/godefault/caller/zap"
)

var cfg = `
[callerGinSession]
	name = "mysession"
    size = 10
    network = "tcp"
    addr = "127.0.0.1:6379"
    pwd = ""
    keypairs = "secret"
[callerGorm.default]
    debug = true
    level = "panic"
    network = "tcp"
    dialect = "mysql"
    addr = "127.0.0.1:3306"
    username = "root"
    password = ""
    db = "default"
    charset = "utf8"
    parseTime = "True"
    loc = "Local"
    timeout = "1s"
    readTimeout = "1s"
    writeTimeout = "1s"
    maxOpenConns = 30
    maxIdleConns = 10
    connMaxLifetime = "300s"
[callerZap.default]
    debug = true
    level = "debug"
    path = "./system.log"

`
var (
	Db      *gorm.GormClient
	Logger  *zap.ZapClient
	Session gin.HandlerFunc
)

func main() {
	caller.Init(
		cfg,
		zap.New,
		gorm.New,
		ginsession.New,
	)

	initModel()

	type User struct {
		Uid  int
		Name string
	}

	u := User{}
	Db.Table("user").Where("uid=?", 1).Find(&u)

	Logger.Info("hello world")
	r := gin.New()

	r.Use(model.Session)

}

func initModel() {
	Db = gorm.Caller("default")
	Logger = zap.Caller("default")
	Session = ginsession.Caller()
}
