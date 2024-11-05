package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/helper"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
	this middleware is for logging request and response
	this logger will write to file ./logs/app_{date}.log every day
	you can use this middleware in your route

	how to use it
	1. import middleware
	2. use middleware.LogMiddleware

	more info contact me @marifsulaksono
*/

type CustomResponseWriter struct {
	echo.Response
	Body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b) // capture the response body
	return w.Response.Writer.Write(b)
}

type LogEntry struct {
	URL      string      `json:"url"`
	Method   string      `json:"method"`
	IP       string      `json:"ip"`
	User     interface{} `json:"user"`
	Body     interface{} `json:"body"`
	Response interface{} `json:"response"`
}

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload, err := helper.GetPayloadAndRecycle(c)
		if err != nil {
			return response.BuildErrorResponse(c, err)
		}
		logger := initLogger()

		res := &CustomResponseWriter{
			Response: *c.Response(),
			Body:     new(bytes.Buffer),
		}
		c.Response().Writer = res
		err = next(c)

		// log the response
		var response interface{}
		if err == nil {
			json.Unmarshal(res.Body.Bytes(), &response)
		}

		entry := LogEntry{
			URL:      c.Request().URL.String(),
			Method:   c.Request().Method,
			IP:       c.RealIP(),
			Body:     payload,
			Response: response,
		}

		logger.WithFields(logrus.Fields{
			"log": entry,
		}).Info("Request log")

		return err
	}
}

// init new logger
func initLogger() *logrus.Logger {
	currentLogDate := time.Now().Format("20060102")
	filename := fmt.Sprintf("./logs/app_%s.log", currentLogDate)

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	logger.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,   // Maximum size in MB before rotation
		MaxBackups: 7,    // Number of backups to keep
		MaxAge:     30,   // Maximum age of logs in days
		Compress:   true, // Compress old log files
	})

	return logger
}
