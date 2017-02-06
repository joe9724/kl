package main

import (
	_ "kl/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)


func main() {
	fmt.Println("init mysql ")
	orm.RegisterDriver("mysql",orm.DRMySQL);
	orm.RegisterDataBase("default","mysql","pf:123456@tcp(192.168.30.103:63306)/ibus?charset=utf8")
	o := orm.NewOrm()
	o.Using("default")
	var count int
	o.Raw("select count(*) as count from news").QueryRow(&count)
	fmt.Println(count)
	//
	beego.Run()



}

