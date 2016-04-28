package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-oauth2/serversample/routers"
)

func main() {
	beego.Run()
}
