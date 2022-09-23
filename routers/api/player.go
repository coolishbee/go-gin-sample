package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jameschun7/go-gin-sample/pkg/app"
	"github.com/jameschun7/go-gin-sample/pkg/e"
	"github.com/jameschun7/go-gin-sample/service/player_service"
	"github.com/unknwon/com"
)

// @Summary Get player list by team id
// @Produce  json
// @Param team_id path int true "TEAM ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/playerlist/{team_id} [get]
func GetPlayerList(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("team_id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "team_id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(e.INVALID_PARAMS, nil)
		return
	}

	playerService := player_service.Team{
		TeamID: id,
	}
	playerList, err := playerService.GetPlayerList()
	if err != nil {
		appG.Response(e.ERROR_DB, nil)
		return
	}
	data := make(map[string]interface{})
	data["team"] = playerList
	appG.Response(e.SUCCESS, data)
}
