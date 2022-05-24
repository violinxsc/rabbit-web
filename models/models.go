package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int    `form:"-"`
	Name    string `form:"name"`
	Sex     string `form:"sex"`
	Remarks string `form:"remarks"`
}

func init() {
	logs.Info("init db config")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	datasource, _ := web.AppConfig.String("mysqldatasource")
	orm.RegisterDataBase("default", "mysql", datasource)
	orm.RegisterModel(new(User))
}
