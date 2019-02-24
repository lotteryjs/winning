package controller

import (
	"github.com/lotteryjs/winning/util"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/lotteryjs/winning/log"
	"github.com/lotteryjs/winning/model"
)

// Logger
var logger = log.NewLogger(os.Stdout)

// MapRoutes returns a gin engine and binds controllers with request URLs.
func MapRoutes() *gin.Engine {
	ret := gin.New()

	if "dev" == model.Conf.RuntimeMode {
		ret.Use(gin.Logger())
	}
	ret.Use(gin.Recovery())

	store := cookie.NewStore([]byte(model.Conf.SessionSecret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   model.Conf.SessionMaxAge,
		Secure:   strings.HasPrefix(model.Conf.Server, "https"),
		HttpOnly: true,
	})
	ret.Use(sessions.Sessions("winning", store))

	api := ret.Group(util.PathAPI)
	api.GET("/wnn", func(c *gin.Context) {
		c.String(200, "Hello %s", "WINNING")
	})

	ret.NoRoute(func(c *gin.Context) {
		notFound(c)
	})

	return ret
}
