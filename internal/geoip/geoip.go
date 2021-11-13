package geoip

import (
	// "fmt"
	"net"
	// "net/http"

	"github.com/midoks/go-p2p-server/internal/conf"
	"github.com/oschwald/geoip2-golang"
)

var geoIp *geoip2.Reader

func Init() {
	geoIp, _ = geoip2.Open(conf.Geo.Path)
}

func GetLatLongByIpAddr(ipAddr string) (float64, float64) {
	ip := net.ParseIP(ipAddr)

	// fmt.Println("ip:", ip)
	record, _ := geoIp.City(ip)
	return record.Location.Latitude, record.Location.Longitude
}
