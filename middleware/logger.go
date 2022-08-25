package middleware

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	// 定义日志输出到文件
	filePath := "log/goblog.log" // 必须指定一个文件，否则报错
	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	logger := logrus.New()
	logger.Out = scr

	// 定义日志分割
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotatelogs.New(
		// 定义分割文件名称
		strings.Split(filePath, ".")[0]+"_%Y%m%d.log",
		// 定义日志有效期和分割策略
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		//
		rotatelogs.WithLinkName(filePath),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02:15:04:05",
	})
	logger.AddHook(Hook)

	// 定义日志内容格式
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := int(math.Ceil(float64(stopTime.Nanoseconds()) / 1000000.0))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.ClientIP()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
