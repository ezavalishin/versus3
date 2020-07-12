package transformations

import (
	gql "github.com/ezavalishin/versus3/internal/gql/models"
	dbm "github.com/ezavalishin/versus3/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBRoundToGQLRound(i *dbm.Round) (o *gql.Round, err error) {
	o = &gql.Round{
		ID:   i.ID,
		Step: i.Step,
	}
	return o, err
}
