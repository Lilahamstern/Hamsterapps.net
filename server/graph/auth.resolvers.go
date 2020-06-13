package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lilahamstern/hamsterapps.net/server/graph/generated"
	"github.com/lilahamstern/hamsterapps.net/server/graph/model"
	"github.com/lilahamstern/hamsterapps.net/server/internal/auth"
	"github.com/lilahamstern/hamsterapps.net/server/internal/users"
	"github.com/lilahamstern/hamsterapps.net/server/pkg/jwt"
)

func (r *mutationResolver) Signup(ctx context.Context, input model.SignupInput) (string, error) {
	var user users.User

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password

	user.Create()

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input *model.LoginInput) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password

	if !user.Authenticate() {
		return "", &users.IncorrectEmailOrPasswordError{}
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input *model.RefreshTokenInput) (string, error) {
	user, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", &auth.AccessDeniedError{}
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
