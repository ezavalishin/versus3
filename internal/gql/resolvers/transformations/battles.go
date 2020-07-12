package transformations

import (
	gql "github.com/ezavalishin/versus3/internal/gql/models"
	dbm "github.com/ezavalishin/versus3/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBBattleToGQLBattle(i *dbm.Battle) (o *gql.Battle, err error) {
	o = &gql.Battle{
		ID:    i.ID,
		Title: i.Title,
	}
	return o, err
}
