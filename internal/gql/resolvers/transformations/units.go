package transformations

import (
	gql "github.com/ezavalishin/versus3/internal/gql/models"
	dbm "github.com/ezavalishin/versus3/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUnitToGQLUnit(i *dbm.Unit) (o *gql.Unit, err error) {
	o = &gql.Unit{
		ID:    i.ID,
		Title: i.Title,
	}
	return o, err
}
