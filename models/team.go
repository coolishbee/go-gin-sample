package models

import "github.com/jinzhu/gorm"

type Team struct {
	ID          int    `gorm:"primary_key" json:"id"`
	TeamName    string `json:"team_name"`
	ManagerName string `json:"manager_name"`
	Formation   string `json:"formation"`
}

func GetAllTeamList() ([]*Team, error) {
	var teams []*Team
	err := db.Find(&teams).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return teams, nil
}

func AddTeam(teamName, managerName, formation string) error {
	team := Team{
		TeamName:    teamName,
		ManagerName: managerName,
		Formation:   formation,
	}
	if err := db.Create(&team).Error; err != nil {
		return err
	}
	return nil
}
