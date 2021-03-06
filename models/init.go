package models

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var o orm.Ormer

func ormSetup() {
	dbHost := beego.AppConfig.String("db.host")
	dbType := beego.AppConfig.String("db.type")
	dbName := beego.AppConfig.String("db.name")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbTimezone := beego.AppConfig.String("db.timezone")
	dbCharset := beego.AppConfig.String("db.charset")
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbCharset == "" {
		dbCharset = "utf8"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset)
	if dbTimezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(dbTimezone)
	} else {
		dsn = dsn + "&loc=Local"
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", dbType, dsn)
	orm.RegisterModel(
		new(User),
		new(Job),
		new(JobState),
	)
	if beego.AppConfig.String("runmode") != "prod" {
		orm.Debug = true
	}

	o = orm.NewOrm()
}

func GetOrm() orm.Ormer {
	return o
}

func init() {
	ormSetup()
}
