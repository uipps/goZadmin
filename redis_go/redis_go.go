package redis_go

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Redis_go() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis success!", err)
	}
	defer c.Close()
}