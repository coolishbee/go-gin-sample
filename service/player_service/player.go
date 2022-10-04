package player_service

import "github.com/coolishbee/go-gin-sample/models"

type Team struct {
	TeamID int
}

func (team *Team) GetPlayerList() ([]*models.Player, error) {
	return models.GetPlayerListByTeamID(team.TeamID)
}
