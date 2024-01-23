package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"time"
	"ulog-backend/biz/handler"
	"ulog-backend/biz/model"
	"ulog-backend/config"
)

var (
	AuthMiddleware *jwt.HertzJWTMiddleware
)

func JWTInit() {
	conf := config.JWT
	var err error
	// the auth middleware
	AuthMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:        "ulog",
		Key:          []byte(conf.Key),
		Timeout:      12 * time.Hour,
		MaxRefresh:   12 * time.Hour,
		IdentityKey:  conf.IdentityKey,
		TokenLookup:  "cookie:user_token",
		SendCookie:   true,
		CookieMaxAge: 12 * time.Hour,
		CookieName:   "user_token",
		PayloadFunc: func(data any) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					conf.IdentityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) any {
			claims := jwt.ExtractClaims(ctx, c)
			return &model.User{
				Name: claims[conf.IdentityKey].(string),
			}
		},
		Authenticator: handler.AuthenticateUser,
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]any{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic("JWT Error:" + err.Error())
	}
}
