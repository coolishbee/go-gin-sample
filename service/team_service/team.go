package team_service

import "github.com/jameschun7/go-gin-sample/models"

func GetAllTeamList() ([]*models.Team, error) {
	return models.GetAllTeamList()
}
