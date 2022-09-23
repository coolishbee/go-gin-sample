package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jameschun7/go-gin-sample/pkg/app"
	"github.com/jameschun7/go-gin-sample/pkg/e"
	"github.com/jameschun7/go-gin-sample/service/team_service"
)

// @Summary Get team list
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/teamlist [get]
func GetTeamList(c *gin.Context) {
	appG := app.Gin{C: c}

	teamList, err := team_service.GetAllTeamList()
	if err != nil {
		appG.Response(e.ERROR_DB, nil)
		return
	}
	data := make(map[string]interface{})
	data["list"] = teamList
	appG.Response(e.SUCCESS, data)
}
