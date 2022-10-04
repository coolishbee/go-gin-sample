package api

import (
	"github.com/coolishbee/go-gin-sample/pkg/app"
	"github.com/coolishbee/go-gin-sample/pkg/e"
	"github.com/coolishbee/go-gin-sample/service/team_service"
	"github.com/gin-gonic/gin"
)

type TeamInfo struct {
	TeamName    string `json:"team_name" binding:"required"`
	ManagerName string `json:"manager_name" binding:"required"`
	Formation   string `json:"formation" binding:"required"`
}

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

// @Summary Add TeamInfo
// @Produce  json
// @Param team body TeamInfo true "TeamInfo Body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/team [post]
func AddTeamInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	teamInfo := &TeamInfo{}

	errCode := app.BindAndValid(c, teamInfo)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}
	teamService := team_service.Team{
		TeamName:    teamInfo.TeamName,
		ManagerName: teamInfo.ManagerName,
		Formation:   teamInfo.Formation,
	}
	err := teamService.AddTeamInfo()
	if err != nil {
		appG.Response(e.ERROR_DB, nil)
		return
	} else {
		appG.Response(e.SUCCESS, nil)
	}

}
