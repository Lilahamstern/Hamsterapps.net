package projects

import (
	"github.com/google/uuid"
	database "github.com/lilahamstern/hamsterapps.net/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/hamsterapps.net/server/internal/users"
	"log"
)

type Project struct {
	ID          uuid.UUID   `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	User        *users.User `json:"user"`
}

func (project Project) Save() uuid.UUID {
	query := "INSERT INTO hamsterapps.public.projects(title,description,user_id) VALUES($1,$2,$3) RETURNING id"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(project.Title, project.Description, project.User.ID).Scan(&project.ID)
	if err != nil {
		log.Fatal(err)
	}

	return project.ID
}

func GetAll() []Project {
	query := "SELECT p.id, p.title, p.description, u.id, u.username, u.email FROM hamsterapps.public.projects p INNER JOIN hamsterapps.public.users u on p.user_id = u.id"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	var projects []Project
	for rows.Next() {
		var project Project
		var user users.User
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &user.ID, &user.Username, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		project.User = &user
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return projects
}
