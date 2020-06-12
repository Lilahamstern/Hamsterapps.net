package projects

import (
	database "github.com/lilahamstern/hamsterapps.net/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/hamsterapps.net/server/internal/users"
	"log"
)

type Project struct {
	ID          int64
	Title       string
	Description string
	User        *users.User
}

func (project Project) Save() int64 {
	stmt := "INSERT INTO hamsterapps.public.projects(title,description,user_id) VALUES($1,$2,$3) RETURNING id"
	err := database.DB.QueryRow(stmt, project.Title, project.Description, 1).Scan(&project.ID)
	if err != nil {
		log.Fatal(err)
	}

	return project.ID
}
