/*
 * @Author: dzw
 * @Date: 2020-04-05 12:22:01
 * @Last Modified by:   dzw
 * @Last Modified time: 2020-04-05 12:22:01
 */

package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	rdb, err := initRedisClient()
	if err != nil {
		fmt.Println("link redis server failed,", err)
		return
	}

	// set
	err = rdb.Set("name", "dzw", 0).Err()
	if err != nil {
		fmt.Println("set name failed,", err)
		return
	}

	// set timeout
	err = rdb.Set("login", "dzwwww", 10*time.Second).Err() // 10s
	if err != nil {
		fmt.Println("set login failed,", err)
		return
	}

	// get
	val, err := rdb.Get("name").Result()
	if err != nil {
		fmt.Println("get name failed,", err)
		return
	}
	fmt.Println("get name:", val)

	val, err = rdb.Get("test").Result()
	if err != nil {
		fmt.Println("get test failed,", err)
		return
	}
	fmt.Println("get test:", val)

	val, err = rdb.Get("dzw").Result()
	if err != nil {
		fmt.Println("get dzw failed,", err)
		return
	}
	fmt.Println("get dzw:", val)

	// get not exist key
	// val, err = rdb.Get("no").Result()
	// if err != nil {
	// 	fmt.Println("get no failed,", err)
	// 	return
	// }
	// fmt.Println("get no:", val)

	//zset
	zsetkey := "language_rank"
	languages := []*redis.Z{
		&redis.Z{Score: 90, Member: "Golang"},
		&redis.Z{Score: 98, Member: "Java"},
		&redis.Z{Score: 95, Member: "Python"},
		&redis.Z{Score: 97, Member: "JS"},
		&redis.Z{Score: 99, Member: "C/C++"},
	}
	// ZAdd
	num, err := rdb.ZAdd(zsetkey, languages...).Result()
	if err != nil {
		fmt.Println("zadd failed,", err)
		return
	}
	fmt.Println("zadd", num, "success")

	//golang score add 10
	newScore, err := rdb.ZIncrBy(zsetkey, 10, "Golang").Result()
	if err != nil {
		fmt.Println("zincr failed,", err)
		return
	}
	fmt.Println("Golang's score is", newScore)

	// get top 3
	ret, err := rdb.ZRevRangeWithScores(zsetkey, 0, 2).Result()
	if err != nil {
		fmt.Println("zrevrangewithscore failed,", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// get score 95-100
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}

	ret, err = rdb.ZRangeByScoreWithScores(zsetkey, op).Result()
	if err != nil {
		fmt.Println("zrangebyscorewithscores failed", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func initRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // default
	})

	s, err := rdb.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("link ping result:", s)
	return rdb, nil
}
