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
	var resProjects []*model.Project
	dbProjects := projects.GetAll()

	for _, project := range dbProjects {
		resProjects = append(resProjects, &model.Project{
			ID:          strconv.FormatInt(project.ID, 10),
			Title:       project.Title,
			Description: project.Description,
			User:        nil,
		})
	}

	return resProjects, nil
}
