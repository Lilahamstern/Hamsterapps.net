package projects

import (
	database "github.com/lilahamstern/hamsterapps.net/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/hamsterapps.net/server/internal/users"
	"log"
)

type Project struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	User        *users.User `json:"user"`
}

func (project Project) Save() int64 {
	sql := "INSERT INTO hamsterapps.public.projects(title,description,user_id) VALUES($1,$2,$3) RETURNING id"

	stmt, err := database.DB.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow().Scan(&project.ID)
	if err != nil {
		log.Fatal(err)
	}

	return project.ID
}

func GetAll() []Project {
	sql := "SELECT id, title, description FROM hamsterapps.public.projects"

	stmt, err := database.DB.Prepare(sql)
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
		err := rows.Scan(&project.ID, &project.Title, &project.Description)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return projects
}
