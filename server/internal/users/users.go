package users

import (
	"database/sql"
	database "github.com/lilahamstern/hamsterapps.net/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/hamsterapps.net/server/pkg/hash"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Create() {
	query := "INSERT INTO hamsterapps.public.users(username, email, password) VALUES($1,$2,$3)"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.Username, user.Email, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func (user *User) Authenticate() bool {
	query := "SELECT password FROM users WHERE email = $1"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	var hashedPassword string
	err = stmt.QueryRow(user.Email).Scan(&hashedPassword)
	if err != nil {

		if err == sql.ErrNoRows {
			return false
		}
		log.Fatal(err)
		return false
	}

	return hash.CheckPasswordHash(user.Password, hashedPassword)
}

func GetUserByEmail(email string) (User, error) {
	query := "SELECT id, email, username  FROM users WHERE email = $1"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var user User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		return User{}, err
	}

	return user, nil
}
