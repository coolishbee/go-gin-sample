package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	_ "github.com/coolishbee/go-gin-sample/docs"
	"github.com/coolishbee/go-gin-sample/pkg/setting"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/coolishbee/go-gin-sample/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	store, _ := redis.NewStore(
		setting.RedisSetting.MaxIdle,
		"tcp",
		setting.RedisSetting.Host,
		setting.RedisSetting.Password,
		[]byte("secret"))

	r.Use(sessions.Sessions("gamepub", store))

	r.POST("/api/login", api.Login)
	r.GET("/api/autologin", api.AutoLogin)
	r.GET("/api/teamlist", api.GetTeamList)
	r.GET("/api/playerlist/:team_id", api.GetPlayerList)
	r.POST("/api/team", api.AddTeamInfo)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
