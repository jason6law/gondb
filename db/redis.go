package db

import (
	"github.com/go-redis/redis"
	"fmt"
)

var Client *redis.Client
var address = ""
var password = ""
var db = 0

func init() {

	opt := &redis.Options{
		Addr:address,
		Password:password,
		DB:db,
	}
	Client = redis.NewClient(opt)
	pong,err := Client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("connect to redis error:%v\n",err))
	}
	fmt.Println("redis: ",pong)

}

func Close() {
	Client.Close()
}
