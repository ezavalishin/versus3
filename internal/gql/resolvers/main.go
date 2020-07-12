package resolvers

import (
	"github.com/ezavalishin/versus3/internal/gql"
	"github.com/ezavalishin/versus3/internal/orm"
)

// Resolver is a modifable struct that can be used to pass on properties used
// in the resolvers, such as DB access
type Resolver struct {
	ORM *orm.ORM
}

// Query exposes query methods
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
