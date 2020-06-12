package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/lilahamstern/hamsterapps.net/server/graph/model"
	"github.com/lilahamstern/hamsterapps.net/server/internal/projects"
	"strconv"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	var project projects.Project
	project.Title = input.Title
	project.Description = input.Description
	projectId := project.Save()

	return &model.Project{
		ID:          strconv.FormatInt(projectId, 10),
		Title:       project.Title,
		Description: project.Description,
		User:        nil,
	}, nil
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
