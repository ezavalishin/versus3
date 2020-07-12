package resolvers

import (
	"context"
	"github.com/ezavalishin/versus3/internal/auth"
	"github.com/ezavalishin/versus3/internal/gql/models"
	"github.com/ezavalishin/versus3/internal/gql/resolvers/transformations"
)

func (r *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {

	return authedUser(r, ctx)
}

func authedUser(r *queryResolver, ctx context.Context) (*models.User, error) {

	user := auth.ForContext(ctx)

	return transformations.DBUserToGQLUser(user)
}
