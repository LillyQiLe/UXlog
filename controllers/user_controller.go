package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"uxlog/models"
	"github.com/pkg/errors"
	"time"
	"crypto/md5"
	"encoding/hex"
	"path"
	"strings"
	"strconv"
	"crypto"
	"net/http"
)


//register controller
type RegisterController struct {
	beego.Controller
}

//CheckUserNameController
type CheckUserNameController struct {
	beego.Controller
}

//LoginController
type LoginController struct {
	beego.Controller
}

//AvatarController
type AvatarController struct {
	beego.Controller
}

//POST: create user
//json
/*
{
    "UserName": "leo",
    "UserSex": 3,
    "UserEmail": "thinkelreo@hotmail.com",
    "PicUrl": "http://127.0.0.1/avatar/thinkerleo.jpg"
}
*/

func (this *RegisterController) Post() {
	type user_info struct{
		UserName string		`form:"UserName"`
		Password string		`form:"Password"`
		UserSex int			`form:"UserSex"`
		UserEmail string	`form:"UserEmail"`
		PicUrl string		`form:"PicUrl"`
	}

	type res_info struct{
		StateCode int
	}

	var res res_info
	var user user_info

	res.StateCode = 1

	if err := this.ParseForm(&user); err != nil {
		res.StateCode = 0
	}

	if _,err := check_password(user.Password); err != nil{
		res.StateCode = 0;
	}

	user.Password = get_md5(user.Password)

	if _, err := check_sex(user.UserSex); err!=nil{
		res.StateCode = 0
	}

	if _, err := check_email(user.UserEmail); err != nil{
		res.StateCode = 0
	}

	if(res.StateCode == 1){
		db := models.DBtools{}
		db.InitDataBase()
		db.OpenDB()
		reg_time := time.Now().Unix()
		if _, err:= db.InsertUser(user.UserName, user.Password, user.UserSex, user.UserEmail, reg_time,user.PicUrl); err !=nil{
			println(err.Error())
			res.StateCode = 0
		}else{
			res.StateCode = 1
		}
	}


	if res_json, err := json.Marshal(res); err != nil{
		this.Ctx.Abort(500, "error")
	}else {
		this.Data["json"] = string(res_json)
		this.ServeJSON()
	}
}

func (this *RegisterController)Get(){
	userName := this.GetString("UserName")
	v := this.GetSession("login_state")
	if v == nil {
		this.Data["json"] = `{"StateCode":0}`
		this.ServeJSON()
	} else {
		if(v != userName){
			this.Data["json"] = `{"StateCode":0}`
			this.ServeJSON()
		}
	}
	type info struct {
		StateCode   int
		UserId      int
		BlogNum     int
		FansCount   int
		FollowCount int
		PicUrl      string
		UserSex     int
		UserEmail   string
		RegTime     string
	}

	var res info

	if err := get_user_info(userName,&res.UserId, &res.BlogNum, &res.FansCount, &res.FollowCount,
							&res.PicUrl, &res.UserSex, &res.UserEmail, &res.RegTime);
		err != nil{
		this.Abort("404")
	}else{
		res.StateCode = 1;
		if res_json, err := json.Marshal(res); err != nil{
			this.Abort("500")
		}else {
			this.Data["json"] = string(res_json)
		}
	}
	this.ServeJSON()
}

//GET
func (this *CheckUserNameController) Get(){
	userName := this.GetString("UserName")
	if res, err := check_username_valid(userName); err != nil{
		this.Ctx.Abort(500, "error")
	}else{
		if(res == true){
			this.Data["json"] = `{"StateCode":1}`
		}else {
			this.Data["json"] = `{"StateCode":0}`
		}
		this.ServeJSON()
	}

}


//POST
func (this *LoginController) Post(){
	userName := this.GetString("UserName")
	password := this.GetString("Password")

	password = get_md5(password)

	if res, err := verify_user(userName, password); err != nil{
		this.Ctx.Abort(500, "error")
	}else{
		if(res == true){
			v := this.GetSession("login_state")
			if v == nil {
				this.SetSession("login_state", userName)
			} else {
				this.SetSession("login_state", userName)
			}
			this.Data["json"] = `{"StateCode":1}`
		}else {
			this.Data["json"] = `{"StateCode":0}`
		}
		this.ServeJSON()
	}
}

func (this *LoginController) Get(){
	v := this.GetSession("login_state")
	if v == nil {
		this.Data["json"] = `{"StateCode":0}`
	} else {
		var userName string
		userName = v.(string)
		this.Data["json"] = `{"StateCode":1, "UserName":"` + userName +`"}`
	}
	this.ServeJSON()
}


func (this *LoginController) Delete(){
	v := this.GetSession("login_state")
	if v != nil {
		this.DelSession("login_state")
		this.Data["json"] = `{"StateCode":1}`
	}else{
		this.Data["json"] = `{"StateCode":0}`
	}
	this.ServeJSON()
}

func (this *AvatarController) Post(){
	//image，这是一个key值，对应的是html中input type-‘file’的name属性值
	v := this.GetSession("login_state")
	if v == nil {
		this.Abort("404")
	}
	type res struct {
		StateCode int
		URL 	string
	}

	f, h, _ := this.GetFile("image")
	if f == nil{
		this.Abort("404")
	}
	pram := strings.Split(h.Filename, ".")
	img_type := pram[len(pram) - 1]
	//得到文件的名称
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(v.(string) +  strconv.FormatInt(time.Now().Unix(), 10)))
	md5FileName:= md5Ctx.Sum(nil)
	fileName := hex.EncodeToString(md5FileName)
	fileName += "." + img_type
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	//保存文件到指定的位置
	var r res
	if err := this.SaveToFile("image", path.Join("static/avatar",fileName));err != nil {
		r.StateCode = 0
		r.URL = ""
		this.Data["json"] = &r
	}else{
		r.StateCode = 1
		r.URL = "user_avatar/" + fileName
		this.Data["json"] = &r
	}
	this.ServeJSON()
}

func (this *AvatarController) Get(){
	filename := this.Ctx.Input.Param(":filename");
	if(filename == ""){
		this.Abort("404")
	}
	//this.Ctx.Output.Download("static/avatar/" + filename,filename)
	http.ServeFile(this.Ctx.ResponseWriter, this.Ctx.Request, "static/avatar/" + filename)
}




/*-------------------------------------------function-------------------------------------------*/
func check_username_exist(username string)(bool, error){
	 db := models.DBtools{}
	 db.InitDataBase()
	 db.OpenDB()
	if res, err := db.QueryUserName(username); err != nil{
		return false, errors.New("server query database error")
	}else {
		if(res){
			return true, nil
		}else{
			return false, nil
		}
	}
 }

func verify_user(userName string, password string)(bool, error){
	if(userName == "" || password == ""){
		return false, nil
	}else{
		db := models.DBtools{}
		db.InitDataBase()
		db.OpenDB()
		if res, err := db.CheckUser(userName, password); err !=nil{
			return false, err
		}else {
			if(res == true){
				return true, nil
			}else{
				return false, nil
			}
		}
	}
}

func check_username_valid(userName string)(bool, error){
	if(userName == ""){
		return false, nil
	}else{
		if res, err := check_username_exist(userName); err !=nil{
			return false, err
		}else {
			if(res == true){
				return false, nil
			}else{
				return true, nil
			}
		}
	}
}

func get_user_info(userName string,user_id *int,blog_num *int,fans_count* int,follow_count* int,pic_url *string,user_sex *int, user_email *string, reg_time *string)(error){
	db := models.DBtools{}
	db.InitDataBase()
	db.OpenDB()
	if err := db.QueryUserInfo(userName,user_id, blog_num, fans_count, follow_count, pic_url, user_sex, user_email, reg_time); err !=nil{
		return err
	}else {
		return nil
	}
}

 func check_sex(UserSex int)(int , error){
		 if(UserSex == 0 || UserSex == 1 || UserSex == 2 || UserSex == 3){
		 	return UserSex, nil
		 }else {
			 return -1, errors.New("invalid sex")
		 }
 }

func  check_email(UserEmail string) (bool, error) {
	return true, nil
}

func  check_password(UserEmail string) (bool, error) {
	return true, nil
}


func get_md5(str string)(string){
	shaCtx := crypto.SHA256.New()
	shaCtx.Write([]byte(str + sha256_salt))
	shaPasswordHex := shaCtx.Sum(nil)
	return hex.EncodeToString(shaPasswordHex)
}

/*-------------------------------------------const value-------------------------------------------*/
var sha256_salt = "6ae4e8682272f33a6b87c1534e58354869d45807de6c327f5afd7e928db5cc6b"