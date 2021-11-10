package announce

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) ([]string, bool) {

	val, err := rdb.SMembers(key).Result()

	fmt.Println("redis get:", val, err)

	ok, err := rdb.Expire(key, 60*time.Second).Result()
	fmt.Printf("ok=%v err=%v\n", ok, err)

	for _, k := range val {
		_, err := rdb.Get(k).Result()
		if err != nil {
			DelKeyValue(key, k)
		}
	}

	if err != nil {
		return []string{}, false
	}

	return val, true
}

func SetPeerHeartbeat(peer string, expiration time.Duration) (interface{}, error) {
	d, err := rdb.Get(peer).Result()
	if err == nil {
		f, err := rdb.Expire(peer, expiration).Result()
		return f, err
	}
	return d, err
}

func Set(key, val string) {

	_, err := rdb.Set(val, "1", 60*time.Second).Result()
	fmt.Println("redism:", err)

	ret, err := rdb.SAdd(key, val).Result()
	fmt.Println("redis:", ret, err)

	_, err = rdb.Expire(key, 60*time.Second).Result()
	fmt.Printf(" err=%v\n", err)

}

func DelKeyValue(key, val string) bool {
	_, err := rdb.SRem(key, val).Result()
	if err == nil {
		return true
	}
	return false
}

func KeyCount(key string) int64 {
	count, err := rdb.SCard(key).Result()
	if err == nil {
		return count
	}
	return 0
}
