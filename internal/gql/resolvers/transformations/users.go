package transformations

import (
	gql "github.com/ezavalishin/versus3/internal/gql/models"
	dbm "github.com/ezavalishin/versus3/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		ID:        i.ID,
		VkUserID:  i.VkUserId,
		Avatar:    i.Avatar,
		FirstName: i.FirstName,
		LastName:  i.LastName,
	}
	return o, err
}
