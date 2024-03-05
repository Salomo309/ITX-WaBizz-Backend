package repositories

import (
	"database/sql"

	"itx-wabizz/models"
)

// Interface for user repository
type UserRepository interface {
	Insert(user *models.User) error
	GetUserByGoogleID(googleID string) (*models.User, error)
}

// Implementation of user repository
type MySQLUserRepository struct {
	db                    *sql.DB
	insertStmt            *sql.Stmt
	getUserByGoogleIDStmt *sql.Stmt
}

// Function to create new user repository
func NewMySQLUserRepository(db *sql.DB) (*MySQLUserRepository, error) {
	insertStmt, err := db.Prepare("INSERT INTO Users (google_id, email, name, picture, admin) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	getUserByGoogleIDStmt, err := db.Prepare("SELECT user_id, google_id, email, name, picture, admin FROM Users WHERE google_id = ?")
	if err != nil {
		return nil, err
	}

	return &MySQLUserRepository{
		db:                    db,
		insertStmt:            insertStmt,
		getUserByGoogleIDStmt: getUserByGoogleIDStmt,
	}, nil
}

// Function to insert new user into database
func (repo *MySQLUserRepository) Insert(user *models.User) error {
	_, err := repo.insertStmt.Exec(user.Google_ID, user.Email, user.Name, user.Picture, user.Admin)
	if err != nil {
		return err
	}
	return nil
}

// Function to get user based on the Google ID
func (repo *MySQLUserRepository) GetUserByGoogleID(googleID string) (*models.User, error) {
	row := repo.getUserByGoogleIDStmt.QueryRow(googleID)

	var user models.User
	err := row.Scan(&user.User_ID, &user.Google_ID, &user.Email, &user.Name, &user.Picture, &user.Admin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
