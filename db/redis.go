package db

import (
	"github.com/go-redis/redis"
	"fmt"
)


var client *redis.Client
var address = ""
var Password = ""
var db = 0

func init() {

	opt := &redis.Options{
		Addr:address,
		Password:Password,
		DB:db,
	}

	client = redis.NewClient(opt)
	pong,err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("connect to redis error:%v\n",err))
	}
	fmt.Println("redis: ",pong)

}

func Close() {
	client.Close()
}
