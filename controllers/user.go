package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"rabbit-web/models"
)

type UserController struct {
	web.Controller
}

func (c *UserController) Save() {
	var user models.User
	//c.ParseForm(&user)
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	o := orm.NewOrm()
	var queryUser models.User
	queryUser.Id = user.Id
	err := o.Read(&queryUser)
	logs.Info(queryUser)
	if err == orm.ErrNoRows {
		result, _ := o.Insert(&user)
		returnDto := models.SuccessResult(result)
		c.Data["json"] = &returnDto
	} else if err == nil {
		result, _ := o.Update(&user)
		returnDto := models.SuccessResult(result)
		c.Data["json"] = &returnDto
	}
	c.ServeJSON()
}

func (c *UserController) GetById() {
	o := orm.NewOrm()
	var user models.User
	user.Id, _ = c.GetInt("id")
	if err := o.Read(&user); err != nil {
		returnDto := models.FailResult(err)
		c.Data["json"] = &returnDto
	} else {
		returnDto := models.SuccessResult(user)
		c.Data["json"] = &returnDto
	}
	c.ServeJSON()
}

func (c *UserController) DeleteById() {
	o := orm.NewOrm()
	var user models.User
	user.Id, _ = c.GetInt("id")
	if num, err := o.Delete(&user); err != nil {
		returnDto := models.FailResult(err)
		c.Data["json"] = &returnDto
	} else {
		returnDto := models.SuccessResult(num)
		c.Data["json"] = &returnDto
	}
	c.ServeJSON()
}
