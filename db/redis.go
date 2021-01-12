package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis(host string, port int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Printf("redis: %v", rdb)
}

func Pop() {
	for {
		val, err := rdb.BRPop(ctx, 0, "trumpeter:request:line").Result()
		//fmt.Printf("打印: %s \n", val)
		//
		//fmt.Printf("打印2: %s \n", val[1])
		switch {
		case err == redis.Nil:
			fmt.Println("key does not exist")
		case err != nil:
			fmt.Println("BRPop failed", err)
		case val[0] == "":
			fmt.Println("value is empty")
		}

		Save(val[1])
	}
}

func Save(s string) {
	var f interface{}

	err := json.Unmarshal([]byte(s), &f)
	if err != nil {
		fmt.Println(err)
	}

	m := f.(map[string]interface{})

	for k, v := range m {

		fmt.Printf("key: %v value: %v \n", k, v)

		var item interface{}

		err := json.Unmarshal([]byte(s), &item)
		if err != nil {
			fmt.Println(err)
		}

		row := item.(map[string]interface{})

		for k1, v1 := range row {
			fmt.Printf("== k1: %v v1: %v \n", k1, v1)
		}

		//switch vv := v.(type) {
		//case string:
		//	fmt.Println(k, "is string", vv)
		//case int:
		//	fmt.Println(k, "is int", vv)
		//case []interface{}:
		//	fmt.Println(k, "is an array:")
		//	for i, u := range vv {
		//		fmt.Println(i, u)
		//	}
		//default:
		//	fmt.Println(k, "is of a type I don't know how to handle")
		//}
	}
}
