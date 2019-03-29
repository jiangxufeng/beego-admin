package main

import (
	_ "myblog/routers"
	"myblog/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	models.Init()
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}

func main() {

	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//同步 ORM 对象和数据库
	//这时, 在你重启应用的时候, beego 便会自动帮你创建数据库表。
	orm.Debug = true

	orm.RunSyncdb("default", false, true)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Content-Type", "Set-Cookie"},
		AllowCredentials: true,
	}))
	// orm.RunCommand()
	beego.Run()
}

