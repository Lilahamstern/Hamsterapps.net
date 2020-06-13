package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/lilahamstern/hamsterapps.net/server/graph/model"
	"github.com/lilahamstern/hamsterapps.net/server/internal/auth"
	"github.com/lilahamstern/hamsterapps.net/server/internal/projects"
	"log"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	user, err := auth.ForContext(ctx)
	if err != nil {
		log.Printf("Error occuerd while fetching context: %s", err)
		return &model.Project{}, fmt.Errorf("internal server error")
	}

	if user == nil {
		return &model.Project{}, &auth.AccessDeniedError{}
	}

	var project projects.Project
	project.Title = input.Title
	project.Description = input.Description
	project.User = user
	projectId := project.Save()

	return &model.Project{
		ID:          projectId.String(),
		Title:       project.Title,
		Description: project.Description,
		User: &model.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var resProjects []*model.Project
	dbProjects := projects.GetAll()

	for _, project := range dbProjects {
		resProjects = append(resProjects, &model.Project{
			ID:          project.ID.String(),
			Title:       project.Title,
			Description: project.Description,
			User: &model.User{
				ID:       project.User.ID,
				Username: project.User.Username,
				Email:    project.User.Email,
			},
		})
	}

	return resProjects, nil
}
