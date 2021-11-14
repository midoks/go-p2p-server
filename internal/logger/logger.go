package logger

import (
	// "fmt"
	"os"
	"strings"

	"github.com/midoks/go-p2p-server/internal/conf"
	go_logger "github.com/phachon/go-logger"
)

var logger *go_logger.Logger

func Init() {

	jsonFormat := false
	if strings.EqualFold(conf.Log.Format, "json") {
		jsonFormat = true
	}

	logDir := conf.Log.RootPath
	os.MkdirAll(logDir, os.ModePerm)

	logger = go_logger.NewLogger()

	// 文件输出配置
	fileConfig := &go_logger.FileConfig{
		Filename: logDir + "/info.log",
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): logDir + "/error.log", // Error 级别日志被写入 error .log 文件
			logger.LoggerLevel("info"):  logDir + "/info.log",  // Info 级别日志被写入到 info.log 文件中
			logger.LoggerLevel("debug"): logDir + "/debug.log", // Debug 级别日志被写入到 debug.log 文件中
		},
		MaxSize:    1024 * 1024,                                                    // 文件最大值（KB），默认值0不限
		MaxLine:    100000,                                                         // 文件最大行数，默认 0 不限制
		DateSlice:  "d",                                                            // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: jsonFormat,                                                     // 写入文件的数据是否 json 格式化
		Format:     "%millisecond_format% [%level_string%] [%file%:%line%] %body%", // 如果写入文件的数据不 json 格式化，自定义日志格式
	}
	// 添加 file 为 logger 的一个输出
	logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)

	//default attach console, detach console
	logger.Detach("console")

	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true,
		JsonFormat: false,
		Format:     "%millisecond_format% [%level_string%] [%file%:%line%] %body%",
	}

	logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)
}

func GetLogger() *go_logger.Logger {
	return logger
}

func Debug(args string) {
	logger.SetAsync()
	logger.Debug(args)
	logger.Flush()
}

func Info(args string) {
	logger.SetAsync()
	logger.Info(args)
	logger.Flush()
}

func Warn(args string) {
	logger.Warning(args)
}

func Error(args string) {
	logger.Error(args)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
