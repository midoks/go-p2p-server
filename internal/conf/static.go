package conf

import (
// "net/url"
// "os"
)

// CustomConf returns the absolute path of custom configuration file that is used.
var CustomConf string

var (
	App struct {

		// ⚠️ WARNING: Should only be set by the main package (i.e. "gop2p.go").
		Version string `ini:"-"`

		Name      string
		BrandName string
		RunUser   string
		RunMode   string
	}

	Web struct {
		HttpAddr       string `ini:"http_addr"`
		HttpPort       string `ini:"http_port"`
		HttpServerAddr string `ini:"http_server_addr"`
	}

	// log
	Log struct {
		Format   string
		RootPath string
	}

	Redis struct {
		Enable   bool
		Address  string
		Password string
		Bb       int
		Prefix   string
	}

	Mem struct {
		CacheChannelTime int64
	}

	Geo struct {
		Path string
	}
)
