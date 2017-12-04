// Package middlewares provides common middleware handlers.
package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"context"
	"log"
	"github.com/spf13/viper"
)

func SetDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

func Session(engine *gin.Engine, cfg *viper.Viper) gin.HandlerFunc {
	sc := cfg.Sub("session")
	oc := sc.Sub("options")

	secret1 := sc.GetString("cookie_secret1")
	secret2 := sc.GetString("cookie_secret2")
	var store sessions.Store
	if sc.GetBool("redis.enable") {
		rc := sc.Sub("redis")
		redisProto := rc.GetString("protocol")
		redisAddr := rc.GetString("addr")
		redisPassword := rc.GetString("password")
		var err error
		store, err = sessions.NewRedisStore(10, redisProto, redisAddr, redisPassword, []byte(secret1), []byte(secret2))
		if err != nil {
			log.Printf("Redis connection failed: %v", err)
			// Fall back to cookies
			store = sessions.NewCookieStore([]byte(secret1), []byte(secret2))
		}
	} else {
		store = sessions.NewCookieStore([]byte(secret1), []byte(secret2))
	}
	opts := sessions.Options{
		Path:     oc.GetString("path"),
		Domain:   oc.GetString("domain"),
		MaxAge:   oc.GetInt("maxage"),
		Secure:   oc.GetBool("secure"),
		HttpOnly: oc.GetBool("httponly"),
	}
	store.Options(opts)
	return gin.HandlerFunc(sessions.Sessions("gotime-session", store))
}

// MustLogin is a middleware that checks existence of current user.
func MustLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userRowInterface := session.Get("user")
	if userRowInterface == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	} else {
		ctx.Next()
	}
}
