package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
	"golang.org/x/time/rate"
)

/*
	this middleware is for limiting requests to prevent DDOS attack using leaky bucket algorithm

	how to use it
	1. import middleware
	2. use middleware.RateLimitMiddleware(limit, durationInSec)

	more info contact me @marifsulaksono
*/

var leakyBuckets = make(map[string]*rate.Limiter)
var leakyMu sync.Mutex

func getBucketLimiter(ip string, limit, durationInSec int) *rate.Limiter {
	leakyMu.Lock()
	defer leakyMu.Unlock()

	limiter, exists := leakyBuckets[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Every(time.Duration(durationInSec)*time.Second), limit)
		leakyBuckets[ip] = limiter
	}

	return limiter
}

func RateLimitMiddleware(limit, durationInSec int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()
			limiter := getBucketLimiter(ip, limit, durationInSec)

			if limiter.Allow() {
				return next(c)
			}
			err := response.NewCustomError(http.StatusTooManyRequests, "Permintaan ditolak, coba lagi nanti", nil)
			return response.BuildErrorResponse(c, err)
		}
	}
}
