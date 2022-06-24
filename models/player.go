package models

import "github.com/jinzhu/gorm"

type Player struct {
	ID          int    `gorm:"primary_key" json:"id"`
	PlayerName  string `json:"player_name"`
	SquadNumber string `json:"squad_number"`
	TeamNumber  string `json:"team_number"`
	TeamId      int    `json:"team_id"`
	Nation      string `json:"nation"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
	Age         int    `json:"age"`
	Foot        string `json:"foot"`
	Position    string `json:"position"`
	BallControl int    `json:"ball_control"`
	Dribbling   int    `json:"dribbling"`
	Low_pass    int    `json:"low_pass"`
	Lofted_pass int    `json:"lofted_pass"`
	Speed       int    `json:"speed"`
	Jump        int    `json:"jump"`
}

func GetPlayerListByTeamID(teamid int) ([]*Player, error) {
	var player []*Player
	err := db.Find(&player, "team_id = ?", teamid).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return player, nil
}
