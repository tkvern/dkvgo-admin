package services

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Total    int64 `json:"total"`
	Current  int   `json:"current"`
	PageSize int   `josn:"pageSize"`
}

var (
	o           orm.Ormer
	UserService *userService
	JobService  *jobService
)

func GetOrm() orm.Ormer {
	return o
}

func initServices() {
	UserService = &userService{}
	JobService = &jobService{}
}

func Init() {
	o = orm.NewOrm()
	initServices()
}
