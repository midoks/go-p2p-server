package conf

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"

	"github.com/midoks/go-p2p-server/internal/assets/conf"
	"github.com/midoks/go-p2p-server/internal/tools"
)

// Asset is a wrapper for getting conf assets.
func Asset(name string) ([]byte, error) {
	return conf.Asset(name)
}

// AssetDir is a wrapper for getting conf assets.
func AssetDir(name string) ([]string, error) {
	return conf.AssetDir(name)
}

// MustAsset is a wrapper for getting conf assets.
func MustAsset(name string) []byte {
	return conf.MustAsset(name)
}

// File is the configuration object.
var File *ini.File

//Default production profile
func Init() error {
	customConfFn := "custom/conf/app.conf"
	if tools.IsExist(customConfFn) {
		return InitCostomConf(customConfFn)
	} else {
		os.MkdirAll(filepath.Dir(customConfFn), os.ModePerm)

		cfg := ini.Empty()
		if tools.IsFile(customConfFn) {
			// Keeps custom settings if there is already something.
			if err := cfg.Append(customConfFn); err != nil {
				return errors.Wrap(err, "Failed to load custom conf")
			}
		}

		cfg.Section("").Key("app_name").SetValue("GoP2P")
		cfg.Section("").Key("run_mode").SetValue("prod")

		if err := cfg.SaveTo(customConfFn); err != nil {
			return err
		}

		return InitCostomConf(customConfFn)
	}
	return nil
}

func InitCostomConf(customConf string) error {

	File, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, conf.MustAsset("conf/app.conf"))
	if err != nil {
		return errors.Wrap(err, "parse 'conf/app.conf'")
	}

	File.NameMapper = ini.TitleUnderscore

	if customConf == "" {
		customConf = filepath.Join(CustomDir(), "conf", "app.conf")
	} else {
		customConf, err = filepath.Abs(customConf)
		if err != nil {
			return errors.Wrap(err, "get absolute path")
		}
	}

	CustomConf = customConf

	if tools.IsFile(customConf) {
		if err = File.Append(customConf); err != nil {
			return errors.Wrapf(err, "append %q", customConf)
		}
	} else {
		log.Println("Custom config ", customConf, " not found. Ignore this warning if you're running for the first time")
	}

	if err = File.Section(ini.DefaultSection).MapTo(&App); err != nil {
		return errors.Wrap(err, "mapping default section")
	}

	return nil
}
