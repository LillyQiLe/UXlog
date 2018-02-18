package routers

import (
	"uxlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
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

	beego.AutoRouter(&controllers.RegisterController{})



}
