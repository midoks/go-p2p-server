package mem

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
)

var rdb *redis.Client

const PEER_GEO_NAME = "geo_lat_lng"

func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Address,
		Password: conf.Redis.Password, // no password set
		DB:       conf.Redis.Bb,       // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return err
	}

	go check()
	return nil
}

//检查是否失效
func check() {
	for {
		time.Sleep(10 * time.Second)
		//TODO
		// fmt.Println("check..")
	}
}

func getPrefixKey(key string) string {
	return fmt.Sprintf("%s%s", conf.Redis.Prefix, key)
}

//获取同一资源下的Peer
func GetChannel(key string) ([]string, bool) {
	_key := getPrefixKey(key)

	val, err := rdb.SMembers(_key).Result()

	if err != nil {
		return []string{}, false
	}

	_, err = rdb.Expire(_key, 60*time.Second).Result()
	if err != nil {
		return []string{}, false
	}

	//删除过期数据
	for _, k := range val {
		_peer := getPrefixKey(k)
		_, err := rdb.Get(_peer).Result()
		if err != nil {
			DelChannelPeerValue(key, k)
		}
	}

	if err != nil {
		return []string{}, false
	}

	return val, true
}

//设置同一资源下的Peer
func SetChannel(key, peer string) error {
	_key := getPrefixKey(key)
	_peer := getPrefixKey(peer)

	_, err := rdb.Set(_peer, "0,0", 60*time.Second).Result()
	if err != nil {
		return err
	}

	_, err = rdb.SAdd(_key, peer).Result()
	if err != nil {
		return err
	}

	_, err = rdb.Expire(_key, 60*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}

func SetPeerHeartbeat(key string, expiration time.Duration) (interface{}, error) {
	_key := getPrefixKey(key)
	d, err := rdb.Get(_key).Result()
	if err == nil {
		f, err := rdb.Expire(_key, expiration).Result()
		return f, err
	}
	return d, err
}

func SetPeerLatLang(key string, lat, lng float64) error {
	_key := getPrefixKey(key)
	_, err := rdb.Set(_key, fmt.Sprintf("%f,%f", lat, lng), 60*time.Second).Result()
	if err != nil {
		return err
	}
	GetAllGeoValue()
	AddGeo(key, lat, lng)
	return nil
}

func DelPeer(key string) error {
	_key := getPrefixKey(key)
	_, err := rdb.Del(_key).Result()
	return err
}

func DelChannelPeerValue(key, val string) bool {
	_key := getPrefixKey(key)
	_, err := rdb.SRem(_key, val).Result()
	if err == nil {
		return true
	}
	return false
}

func GetChannelCount(key string) int64 {
	_key := getPrefixKey(key)
	count, err := rdb.SCard(_key).Result()
	if err == nil {
		return count
	}
	return 0
}

func GetAllGeoValue() {
	_key := getPrefixKey(PEER_GEO_NAME)
	r, err := rdb.ZRange(_key, 0, -1).Result()

	for k, v := range r {
		fmt.Println("GetAllGeoValue:", k, v)
	}
	fmt.Println("GetAllGeoValue:", r, err)
}

func AddGeo(label string, lat float64, lng float64) error {
	_key := getPrefixKey(PEER_GEO_NAME)
	_, err := rdb.GeoAdd(_key, &redis.GeoLocation{
		Name:      label,
		Longitude: lng,
		Latitude:  lat,
	}).Result()

	if err != nil {
		logger.Errorf("redis add geo error: %v", err)
		return err
	}
	return nil
}

func DelGeo(label string) error {
	_key := getPrefixKey(PEER_GEO_NAME)
	_, err := rdb.ZRem(_key, label).Result()
	if err != nil {
		logger.Errorf("redis del geo error: %v", err)
		return err
	}
	return nil
}

//存放服务器的经纬度
func GetServerLatLang() (float64, float64, error) {
	key := "server_ip_lat_lng"
	_key := getPrefixKey(key)
	if data, err := rdb.Get(_key).Result(); err == nil {
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

	ip := conf.Web.HttpServerAddr
	lat, lng := geoip.GetLatLongByIpAddr(ip)

	_, err := rdb.Set(_key, fmt.Sprintf("%f,%f", lat, lng), 600*time.Second).Result()
	if err == nil {
		return lat, lng, nil
	}
	return 0, 0, errors.New("not find lat lang")
}
