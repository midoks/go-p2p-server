package announce

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/midoks/go-p2p-server/internal/conf"
	"github.com/midoks/go-p2p-server/internal/geoip"
	"github.com/midoks/go-p2p-server/internal/logger"
	"github.com/midoks/go-p2p-server/internal/tools"
)

var rdb *redis.Client

func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password, // no password set
		DB:       conf.Redis.Bb,       // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		logger.Errorf("redis error: %v", err)
		return err
	}
	return nil
}

func Get(key string) ([]string, bool) {

	val, err := rdb.SMembers(key).Result()

	if err != nil {
		return []string{}, false
	}

	_, err = rdb.Expire(key, 60*time.Second).Result()
	if err != nil {
		return []string{}, false
	}

	//删除过期数据
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

func Set(key, val string) error {

	_, err := rdb.Set(val, "1", 60*time.Second).Result()
	if err != nil {
		return err
	}

	_, err = rdb.SAdd(key, val).Result()
	if err != nil {
		return err
	}

	_, err = rdb.Expire(key, 60*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
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

func GetServerLatLang() (float64, float64, error) {
	key := "go_p2p_server_ip_lat_lng"
	if data, err := rdb.Get(key).Result(); err == nil {
		ll := strings.Split(data, ",")

		lat, err := strconv.ParseFloat(ll[0], 64)
		if err != nil {
			return 0, 0, err
		}

		lng, err := strconv.ParseFloat(ll[1], 64)
		if err != nil {
			return lat, 0, err
		}

		return lat, lng, nil
	}

	ip := tools.GetNetworkIp()
	lat, lng := geoip.GetLatLongByIpAddr(ip)

	_, err := rdb.Set(key, fmt.Sprintf("%f,%f", lat, lng), 600*time.Second).Result()
	if err == nil {
		return lat, lng, nil
	}
	return 0, 0, errors.New("not find lat lang")
}
