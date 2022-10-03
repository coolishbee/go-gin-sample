package team_service

import "github.com/jameschun7/go-gin-sample/models"

type Team struct {
	TeamName    string
	ManagerName string
	Formation   string
}

func GetAllTeamList() ([]*models.Team, error) {
	return models.GetAllTeamList()
}

func (t *Team) AddTeamInfo() error {
	return models.AddTeam(
		t.TeamName,
		t.ManagerName,
		t.Formation)
}
