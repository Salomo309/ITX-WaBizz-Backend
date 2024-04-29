package repositories

import (
	"database/sql"

	"itx-wabizz/models"
)

// Interface for user repository
type UserRepository interface {
	Insert(*models.User) error
	GetUser(string) (*models.User, error)
}

// Implementation of user repository
type MySQLUserRepository struct {
	db                    *sql.DB
	insertStmt            *sql.Stmt
	getUserStmt           *sql.Stmt
}

// Function to create new user repository
func NewMySQLUserRepository(db *sql.DB) (*MySQLUserRepository, error) {
	insertStmt, err := db.Prepare("INSERT INTO Users (email, is_active) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	getUserStmt, err := db.Prepare("SELECT email, is_active FROM Users WHERE email = ?")
	if err != nil {
		return nil, err
	}

	return &MySQLUserRepository{
		db:                    db,
		insertStmt:            insertStmt,
		getUserStmt: 		   getUserStmt,
	}, nil
}

// Function to insert new user into database
func (repo *MySQLUserRepository) Insert(user *models.User) error {
	_, err := repo.insertStmt.Exec(user.Email, user.IsActive)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MySQLUserRepository) GetUser(email string) (*models.User, error) {
	row := repo.getUserStmt.QueryRow(email)

	var user models.User
	err := row.Scan(&user.Email, &user.IsActive)
	if (err != nil) {
		return nil, err
	}

	return &user, nil
}