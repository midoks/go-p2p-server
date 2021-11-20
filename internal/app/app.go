package app

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/midoks/go-p2p-server/internal/assets/public"
	"github.com/midoks/go-p2p-server/internal/assets/templates"
	"github.com/midoks/go-p2p-server/internal/conf"
	"github.com/midoks/go-p2p-server/internal/geoip"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/logger"
	"github.com/midoks/go-p2p-server/internal/mem"
	"github.com/midoks/go-p2p-server/internal/queue"
	// "github.com/midoks/go-p2p-server/internal/tools"
)

const (
	CHECK_CLIENT_INTERVAL = 10
	EXPIRE_LIMIT          = 100
)

func httpCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// }
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %10v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func init() {
	conf.Init()
	logger.Init()
	geoip.Init()

	//App
	queue.Init()
	hub.Init()
	err := mem.Init()
	if err != nil {
		logger.Errorf("init redis error[redis must have]: %v", err)
	}

	go func() {
		for {
			time.Sleep(CHECK_CLIENT_INTERVAL * time.Second)
			now := time.Now().Unix()
			// fmt.Println("start check client alive...")
			count := 0
			for item := range hub.GetInstance().Clients.IterBuffered() {
				cli := item.Val
				if cli.LocalNode && cli.IsExpired(now, EXPIRE_LIMIT) {
					// 节点过期
					logger.Infof("client %s is expired for %d, close it", cli.PeerId, now-cli.Timestamp)
					if ok := hub.DoUnregister(cli.PeerId); ok {
						queue.PushTextLeave(cli.PeerId)
						cli.Close()
						count++
					}
				}
			}
			// fmt.Println("check client finished, closed  clients:", count)
		}
	}()
}

func info(c *gin.Context) {

	m := hub.GetAllKey()
	ipAddr := conf.Web.HttpServerAddr
	lat, lang := geoip.GetLatLongByIpAddr(ipAddr)

	latlng, err := mem.QueryGeoList(m[0], 6)

	c.JSON(http.StatusOK, gin.H{
		"client_num": fmt.Sprintf("%d", hub.GetClientNum()),
		"peers":      m,
		"ip":         ipAddr,
		"lat":        lat,
		"lang":       lang,
		"latlng": gin.H{
			"flatlng", m[0],
			"latlng": latlng,
			"err":    err,
		},
	})
}

func Run() {
	r := gin.New()

	if strings.EqualFold(conf.App.RunMode, "prod") {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(httpCors())
	r.Use(LoggerToFile())
	r.Use(gin.Recovery())

	r.StaticFS("/public", public.BinaryFileSystem(""))
	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		indexByte, _ := templates.Asset("index.tmpl")
		tmpl, _ := template.New("web").Parse(string(indexByte))
		kv := gin.H{"version": conf.App.Version}
		tmpl.Execute(c.Writer, kv)
	})

	r.GET("/getStats", p2pGetStats)
	r.GET("/info", info)

	r.POST("/channel", p2pChannel)
	r.POST("/channel/:channel_id/node/:peer/stats", p2pChannelStats)
	r.POST("/channel/:channel_id/node/:peer/peers", p2pChannelPeers)

	r.GET("/ws", wsSignal)
	r.GET("/ws?id=:id", wsSignal)

	r.Run(fmt.Sprintf(":%s", conf.Web.HttpPort))
}
