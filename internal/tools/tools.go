package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/oschwald/geoip2-golang"
)

var geoIp *geoip2.Reader

func Init() {
	geoIp, _ = geoip2.Open("data/GeoLite2-City.mmdb")
	// defer geoIp.Close()
}

func GetVersionNum(ver string) int {
	digs := strings.Split(ver, ".")
	a, _ := strconv.Atoi(digs[0])
	b, _ := strconv.Atoi(digs[1])
	return a*10 + b
}

type Address struct {
	Country  string `json:"Country"`
	Province string `json:"Province"`
	City     string `json:"City"`
}

type IPLocate struct {
	Result  bool    `json:"result"`
	IP      string  `json:"IP"`
	Address Address `json:"Address"`
	ISP     string  `json:"ISP"`
}

func GetNetworkIp() string {
	conn, err := http.Get("https://ipv4.ipw.cn/api/ip/locate")
	defer conn.Body.Close()
	body, _ := ioutil.ReadAll(conn.Body)
	var ipLocateResult IPLocate
	err = json.Unmarshal(body, &ipLocateResult)
	if err != nil {
		return "127.0.0.1"
	}
	ip := ipLocateResult.IP
	return ip
}

func GetLatLongByIpAddr(ipAddr string) (float64, float64) {
	ip := net.ParseIP(ipAddr)

	fmt.Println("ip:", ip)
	record, _ := geoIp.City(ip)
	return record.Location.Latitude, record.Location.Longitude

}
