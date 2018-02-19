package models

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type DBtools struct{
	dbconn *sql.DB
	dbsetting string
}

var insert_user = `INSERT INTO uxlog.t_user_info (user_name,password, user_sex, user_email, reg_time, pic_url) VALUES (?, ?, ?, ?, ?, ?)`

var cheack_user_name_password = `SELECT t_user_info.user_name FROM uxlog.t_user_info WHERE t_user_info.user_name=? and t_user_info.password=?`

var query_user_name = `SELECT t_user_info.user_name FROM uxlog.t_user_info WHERE t_user_info.user_name=?`

var query_user_info = `SELECT  t_user_info.user_id, t_user_info.blog_num,  t_user_info.fans_count, t_user_info.follow_count, t_user_info.pic_url, t_user_info.user_sex, t_user_info.user_email, t_user_info.reg_time FROM uxlog.t_user_info WHERE t_user_info.user_name=?`

var insert_blog = `INSERT INTO uxlog.t_blog_info (f_user_id, content, type, time) VALUES (?, ?, ?, ?)`

var insert_comment = `INSERT INTO uxlog.t_comment_info (user_id, blog_id, comment_content, time) VALUES (?, ?, ?, ?)`


//Initial database
func (dt*DBtools) InitDataBase(){
	dt.dbconn = new(sql.DB)
	dt.dbsetting = "root:3375@tcp(127.0.0.1:3306)/uxlog?charset=utf8";
}

//Open DataBase
func (dt *DBtools) OpenDB(){
	//Open Databse
	if db, err := sql.Open("mysql", dt.dbsetting);err!=nil{
		panic("Open DataBase Error")
	}else{
		dt.dbconn = db
	}
}

//Insert user
func (dt *DBtools) InsertUser(user_name string, password string, user_sex int, user_email string, reg_time int64, pic_url string)(bool, error){
	dt.OpenDB()
	defer dt.dbconn.Close()
	insert_handle, _ := dt.dbconn.Prepare(insert_user)
	if _, err := insert_handle.Exec(user_name, password, user_sex, user_email, reg_time, pic_url); err!= nil{
		return false, err
	}else {
		return true,nil
	}
}

//InsertBlog
func (dt *DBtools) InsertBlog(user_id uint,content string,blogType int, time int64)(bool, error){
	dt.OpenDB()
	defer dt.dbconn.Close()
	insert_handle, _ := dt.dbconn.Prepare(insert_blog)
	if _, err := insert_handle.Exec(user_id, content, blogType, time); err!= nil{
		return false, err
	}else {
		return true,nil
	}
}

//QueryUserName
func (dt *DBtools) QueryUserName(userName string)(bool, error){
	dt.OpenDB()
	defer  dt.dbconn.Close()
	query_handle, _ := dt.dbconn.Prepare(query_user_name)
	if res, err := query_handle.Query(userName); err != nil{
		return false, err
	}else{
		if res.Next(){
			return true, nil
		}else {
			return false, nil
		}
	}
}

//QueryUserName
func (dt *DBtools) CheckUser(userName string, userPassword string)(bool, error){
	dt.OpenDB()
	defer  dt.dbconn.Close()
	query_handle, _ := dt.dbconn.Prepare(cheack_user_name_password)
	if res, err := query_handle.Query(userName, userPassword); err != nil{
		return false, err
	}else{
		if res.Next(){
			return true, nil
		}else {
			return false, nil
		}
	}
}

//QueryUserInfo
func (dt *DBtools) QueryUserInfo(userName string,user_id *int,blog_num *int,fans_count* int,follow_count* int,pic_url *string,user_sex *int, user_email *string, reg_time *string)(error){
	dt.OpenDB()
	defer  dt.dbconn.Close()
	query_handle, _ := dt.dbconn.Prepare(query_user_info)
	if rows, err := query_handle.Query(userName); err != nil{
		return   err
	}else {
		if rows.Next() {
			err = rows.Scan(user_id, blog_num, fans_count, follow_count, pic_url, user_sex, user_email, reg_time)
			if(err != nil){
				return  err
			}else {
				return nil
			}
		}else {
			return  errors.New("cant find this user")
		}
	}
}
