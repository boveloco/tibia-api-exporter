package models

type CreatureStatistic struct {
	Name  string `json:"race"`
	Count int    `json:"last_day_killed"`
	Date  uint16
}
