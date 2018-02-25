package routers

import (
	"uxlog/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	//AJAX API

	//	Rigister API
	//		POST: create user
	//		PUT: update user information
	//		GET: get user infomation
	beego.Router("/webapi/account", &controllers.RegisterController{})

	//		GET: check whether the user name is occupied
	beego.Router("/webapi/check_username", &controllers.CheckUserNameController{})

	//	Login API
	//		POST: login
	beego.Router("/webapi/login_user", &controllers.LoginController{})
	beego.Router("/webapi/logout", &controllers.LogoutController{})
	beego.Router("/webapi/getinfo", &controllers.GetInfoController{})


	beego.Router("/webapi/getveritycode", &controllers.VerityCodeController{})
	beego.Router("/webapi/checkveritycode", &controllers.VerityCodeController{})



	beego.Router("/webapi/user_avatar/?:filename", &controllers.AvatarController{})


	beego.AutoRouter(&controllers.RegisterController{})



}
