package models

type Unit struct {
	BaseModel
	Battle   Battle
	BattleID int
	Title    string
}
