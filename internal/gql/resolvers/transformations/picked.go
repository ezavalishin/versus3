package transformations

import (
	gql "github.com/ezavalishin/versus3/internal/gql/models"
	dbm "github.com/ezavalishin/versus3/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBPickedToGQLPicked(i *dbm.Picked) (o *gql.Picked, err error) {

	var transformedUsers []*gql.User

	for _, user := range i.Users {
		userGql, _ := DBUserToGQLUser(user)
		transformedUsers = append(transformedUsers, userGql)
	}

	o = &gql.Picked{
		Count:        i.Count,
		FriendsCount: i.FriendsCount,
		Users:        transformedUsers,
	}
	return o, err
}
