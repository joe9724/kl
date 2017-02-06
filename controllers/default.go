package controllers

import (
	"github.com/astaxie/beego"
	"kl/models"
)

type MainController struct {
	beego.Controller
}



func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//
        var list []string
	for i:=0;i<100 ;i++  {
		list = append(list,"abc")
	}
	res := models.Paginator(1, 3, 13)
	c.Data["paginator"] = res


}
