package routers

import (
	"github.com/beego/beego/v2/server/web"
	"rabbit-web/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/user", &controllers.UserController{}, "get:GetById;post:Save;delete:DeleteById")
}
