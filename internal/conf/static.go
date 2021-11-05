package conf

import (
// "net/url"
// "os"
)

var (
	App struct {

		// ⚠️ WARNING: Should only be set by the main package (i.e. "gop2p.go").
		Version string `ini:"-"`

		Name      string
		BrandName string
		RunUser   string
		RunMode   string
	}
)
