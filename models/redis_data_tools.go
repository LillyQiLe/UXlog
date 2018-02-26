package models

import (
	"github.com/garyburd/redigo/redis"
)


type RedisTools struct{
	Network string
	Address string
}
//TODO
//设置过期时间
func (et *RedisTools) AddToken(token_str string, json_str string) (bool, error) {
	c, err := redis.Dial(et.Network, et.Address)
	if err != nil {
		return false, err
	}
	defer c.Close()

	_, err = c.Do("SET", token_str, json_str)
	if err != nil {
		return false, err
	}else {
		return true, nil
	}
}


func (et *RedisTools) GetTokenValue(token_str string)(string, error){
	c, err := redis.Dial(et.Network, et.Address)
	if err != nil {
		return "", err
	}
	defer c.Close()


	res, err := redis.String(c.Do("GET", token_str))
	if err != nil {
		return "", err
	} else {
		return  res, nil
	}

}

func (et *RedisTools) DelToken(token_str string)(error){
	c, err := redis.Dial(et.Network, et.Address)
	if err != nil {
		return err
	}
	defer c.Close()


	_, err = c.Do("DEL", token_str)
	if err != nil {
		return  err
	} else {
		return   nil
	}

}

func (et *RedisTools)PostArticle(user_name string, article_id int32, article_title string, article_type int, likes int64)(bool, error){
	return true, nil
}

func (et *RedisTools)DelArticle(article_id int32)(bool, error){
	return true, nil
}

func (et *RedisTools)LikeArticle(user_name string, article_id int32)(bool, error){
	return true, nil
}

func (et *RedisTools)CancelLikeArticle(user_name string, article_id int32)(bool, error){
	return true, nil
}


