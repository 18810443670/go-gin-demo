package main

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

// 队列的key
var queueKey = "tgvideo_database_queues:doapicalldata"
var rdb *redis.Client

func NewRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pong, "redis success.")
}

// 使用list格式消费消息
func ConsumerMessageList() {
	for {
		// 设置一个5秒的超时时间
		value, err := rdb.BRPop(2*time.Second, queueKey).Result()
		if err == redis.Nil {
			// 查询不到数据
			log.Println("查询不到数据：当前时间是：", time.Now().Unix())
			time.Sleep(1 * time.Second)
			continue
		}
		if err != nil {
			// 查询出错
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("消费到数据：", value[1], "当前时间是：", time.Now().Unix())
		time.Sleep(time.Second)

	}
}

func main() {
	NewRedis()
	ConsumerMessageList()
}
