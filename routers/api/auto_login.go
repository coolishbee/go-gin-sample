package api

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jameschun7/go-gin-sample/pkg/app"
	"github.com/jameschun7/go-gin-sample/pkg/e"
	"github.com/jameschun7/go-gin-sample/service/player_service"
)

// @Summary Get AutoLogin
// @Produce  json
// @Param sessionkey query string true "SessionID for login authentication"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/autologin [get]
func AutoLogin(c *gin.Context) {
	appG := app.Gin{C: c}
	//valid := validation.Validation{}

	playerService := player_service.Team{
		TeamID: 1,
	}

	//sessionkey := c.DefaultQuery("sessionKey", "empty")
	sessionkey := c.Query("sessionkey")
	log.Printf("key: %s", sessionkey)

	session := sessions.Default(c)

	if session.ID() == sessionkey {
		log.Println("세션이 일치합니다")

		userID := session.Get("userID").(string)
		loginType := session.Get("loginType").(string)
		deviceID := session.Get("deviceID").(string)
		socialToken := session.Get("socialToken").(string)

		session.Clear()

		session.Set("userID", userID)
		session.Set("loginType", loginType)
		session.Set("deviceID", deviceID)
		session.Set("socialToken", socialToken)
		session.Save()

		playerList, err := playerService.GetPlayerList()
		if err != nil {
			appG.Response(e.ERROR_DB, nil)
			return
		}
		data := make(map[string]interface{})
		data["sessionKey"] = sessionkey
		data["team"] = playerList
		appG.Response(e.SUCCESS, data)
	} else {
		appG.Response(e.ERROR_DB, nil)
		return
	}
}
