package tools

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
)

func GetVersionNum(ver string) int {
	digs := strings.Split(ver, ".")
	a, _ := strconv.Atoi(digs[0])
	b, _ := strconv.Atoi(digs[1])
	return a*10 + b
}

// IsFile returns true if given path exists as a file (i.e. not a directory).
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !f.IsDir()
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
	if err != nil {
		return "127.0.0.1"
	}

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

func RandId() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var n = 5

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return string(b)
}

// IsExist returns true if a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CurrentUsername returns the username of the current user.
func CurrentUsername() string {
	username := os.Getenv("USER")
	if len(username) > 0 {
		return username
	}

	username = os.Getenv("USERNAME")
	if len(username) > 0 {
		return username
	}

	if user, err := user.Current(); err == nil {
		username = user.Username
	}
	return username
}

// Contains 数组是否包含某元素
func Contains(slice []string, s string) int {
	for index, value := range slice {
		if value == s {
			return index
		}
	}
	return -1
}

func GetOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
