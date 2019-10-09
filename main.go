package main

import (
	"github.com/astaxie/beego"
	_ "github.com/myxzjie/go-cms/routers"
)

func main() {
	// register.
	beego.Run()
}
