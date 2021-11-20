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
	"github.com/midoks/go-p2p-server/internal/tools"
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

		//GEO数据管理
		CheckAllGeoValue()
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

func CheckAllGeoValue() {
	_key := getPrefixKey(PEER_GEO_NAME)
	r, _ := rdb.ZRange(_key, 0, -1).Result()

	//删除过期GEO数据
	for _, k := range r {
		_peer := getPrefixKey(k)
		_, err := rdb.Get(_peer).Result()
		if err != nil {
			err := DelGeo(k)
			fmt.Println("CheckAllGeoValue:", k, _peer, err)
		}
	}
}

//////////////////////////////////////////////////
//https://studygolang.com/articles/27275?fr=sidebar
/////////////////////////////////////////////////
// GEO方法实现

func PosGeo(label string) (float64, float64, error) {
	_key := getPrefixKey(PEER_GEO_NAME)
	resPos, err := rdb.GeoPos(_key, label).Result()

	if err != nil {
		return 0, 0, err
	}

	for _, v := range resPos {
		return v.Latitude, v.Longitude, nil
	}
	return 0, 0, err
}

//获取附近数据|自动重复查找
func QueryGeoList(label string, count int) ([]redis.GeoLocation, error) {
	equatorialCircumference := 40075.02 / 2
	maxQueryTimes := 4
	multiple := 10
	x := tools.GetXForEc(equatorialCircumference, maxQueryTimes, multiple)

	retData := []redis.GeoLocation{}

	for i := 0; i < maxQueryTimes; i++ {
		r, err := QueryGeo(label, tools.GetXForEcIncr(x, i, multiple), count+1)

		if err != nil {
			continue
		}

		if len(r) == 0 {
			continue
		}

		for _, v := range r {
			fmt.Println(v)
			if !strings.EqualFold(label, v.Name) {
				retData = append(retData, v)
			}
		}

		if len(retData) > 0 {
			return retData, nil
		}

		return retData, errors.New("data is empty!")
	}

	return []redis.GeoLocation{}, errors.New("not find geo data!")
}

//获取附近数据
func QueryGeo(label string, dist float64, count int) ([]redis.GeoLocation, error) {
	_key := getPrefixKey(PEER_GEO_NAME)
	lat, lng, err := PosGeo(label)
	if err != nil {
		return []redis.GeoLocation{}, err
	}

	resRadiu, err := rdb.GeoRadius(_key, lng, lat, &redis.GeoRadiusQuery{
		Radius:      dist,  //radius表示范围距离
		Unit:        "km",  //距离单位是 m|km|ft|mi
		WithCoord:   true,  //传入WITHCOORD参数，则返回结果会带上匹配位置的经纬度
		WithDist:    true,  //传入WITHDIST参数，则返回结果会带上匹配位置与给定地理位置的距离
		WithGeoHash: true,  //传入WITHHASH参数，则返回结果会带上匹配位置的hash值
		Count:       count, //入COUNT参数，可以返回指定数量的结果
		Sort:        "ASC", //默认结果是未排序的，传入ASC为从近到远排序，传入DESC为从远到近排序
	}).Result()

	return resRadiu, err
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
