package models

import (
	"testing"
	"time"
	"fmt"
)
var dbsetting = "root:3375@tcp(127.0.0.1:3306)/uxlog?charset=utf8";
func TestDBtools_OpenDB(t *testing.T) {
	db := DBtools{}
	db.InitDataBase(dbsetting)
	db.openDB()
}

func TestDBtools_InsertUser(t *testing.T) {
	db := DBtools{}
	db.InitDataBase(dbsetting)
	db.openDB()
	reg_time := time.Now().Unix()
	res,_ := db.InsertUser(`zzlll`, `78797979`,3, `zll@hotmail.com`,reg_time,`http:\\XXX`)
	println(res)
}

func TestDBtools_InsertBlog(t *testing.T) {
	db := DBtools{}
	db.InitDataBase(dbsetting)
	db.openDB()
	time := time.Now().Unix()
	res, _ := db.InsertBlog(1, "my first blog", 1, time)
	println(res)
}

func TestDBtools_QueryUser(t *testing.T) {
	db := DBtools{}
	db.InitDataBase(dbsetting)
	db.openDB()
	if res, err := db.QueryUserName(`zllll`);  err != nil{
		println(res)
		println(err)
	}else{
		println(res)
	}
}

func TestDBtools_CheckUser(t *testing.T) {
	db := DBtools{}
	db.InitDataBase(dbsetting)
	db.openDB()
	if res, err := db.CheckUser(`zzlll`, "7879797");  err != nil{
		println(res)
		println(err)
	}else{
		println(res)
	}
}

func TestDBtools_QueryUserInfo(t *testing.T) {
	db := DBtools{}
	db.InitDataBase(dbsetting)
	db.openDB()
	var user_id int
	var blog_num int
	var fans_count int
	var follow_count int
	var pic_url string
	var user_sex int
	var user_email string
	var reg_time string
	if err := db.QueryUserInfo("zll",&user_id, &blog_num, &fans_count, &follow_count, &pic_url, &user_sex, &user_email, &reg_time); err !=nil{
	}else {
		fmt.Println(user_id)
		fmt.Println(blog_num)
		fmt.Println(fans_count)
		fmt.Println(user_email)
	}
}