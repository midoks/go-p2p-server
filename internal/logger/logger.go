package logger

import (
	// "fmt"
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

	logger = go_logger.NewLogger()

	// 文件输出配置
	fileConfig := &go_logger.FileConfig{
		Filename: "logs/info.log", // 日志输出文件名，不自动存在
		// 如果要将单独的日志分离为文件，请配置LealFrimeNem参数。
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): "logs/error.log", // Error 级别日志被写入 error .log 文件
			logger.LoggerLevel("info"):  "logs/info.log",  // Info 级别日志被写入到 info.log 文件中
			logger.LoggerLevel("debug"): "logs/debug.log", // Debug 级别日志被写入到 debug.log 文件中
		},
		MaxSize:    1024 * 1024, // 文件最大值（KB），默认值0不限
		MaxLine:    100000,      // 文件最大行数，默认 0 不限制
		DateSlice:  "d",         // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: jsonFormat,  // 写入文件的数据是否 json 格式化
		Format:     "",          // 如果写入文件的数据不 json 格式化，自定义日志格式
	}
	// 添加 file 为 logger 的一个输出
	logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
}

func GetLogger() *go_logger.Logger {
	return logger
}

func Debug(args string) {
	logger.Debug(args)
}

func Info(args string) {
	logger.Info(args)
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
