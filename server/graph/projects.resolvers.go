package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lilahamstern/hamsterapps.net/server/graph/model"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	var project model.Project
	var user model.User
	project.Title = input.Title
	project.Description = input.Description
	user.Email = "leo.ronnebro@hamsterapps.net"
	user.Username = "Lilahamstern"
	project.User = &user

	fmt.Println(ctx)

	return &project, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var projects []*model.Project
	projects = append(projects, &model.Project{
		ID:          "4771723",
		Title:       "Testing",
		Description: "lallala",
		User: &model.User{
			ID:       "1",
			Username: "Lilahamstern",
			Email:    "leo.ronnebro@hamsterapps.net",
		},
	})

	return projects, nil
}
