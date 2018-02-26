package models

import "testing"

func TestAddLoginToken(t *testing.T) {
	conn := RedisTools{"tcp", "127.0.0.1:6379"}
	conn.AddToken("token", "123456")
}

func TestGetLoginStates(t *testing.T) {
	conn := RedisTools{"tcp", "127.0.0.1:6379"}

	res, err := conn.GetTokenValue("dae3e1cc82b97cf40b97f1730ecda15541d2f70a35441c22cd4011f6a9a9f45c")
	if(err != nil){
		println(err.Error())
	}else{
		println(res)
	}
}

func TestRedisTools_DelLoginStates(t *testing.T) {
	conn := RedisTools{"tcp", "127.0.0.1:6379"}
	err := conn.DelToken("3c0d8d679347c49b596ca9b766d5ebd1c8d554ddaa0138f7742e1d64061fd3c7")
	if(err != nil){
		println(err.Error())
	}else{
	}
}