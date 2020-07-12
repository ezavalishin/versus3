package models

type Friend struct {
	BaseModel
	UserOneId int
	UserTwoId int
}
