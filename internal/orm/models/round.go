package models

type Round struct {
	BaseModel
	BattleUser   BattleUser `gorm:"foreignkey:BattleUserId"`
	BattleUserId int
	Step         int
	Pairs        []Pair `gorm:"foreignkey:RoundId"`
}
