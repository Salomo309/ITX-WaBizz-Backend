package repositories

import (
	"database/sql"
	"fmt"

	"itx-wabizz/models"
)

// Interface for user repository
type UserRepository interface {
	Insert(user *models.User) error
	GetUser(email string) (*models.User, error)
	GetAllUser() ([]models.User, error)
	MakeActive(email string) error
	MakeInactive(email string) error
}

// Implementation of user repository
type MySQLUserRepository struct {
	db                    *sql.DB
	insertStmt            *sql.Stmt
	getUserStmt           *sql.Stmt
	getAllUserStmt        *sql.Stmt
	makeActiveStmt		  *sql.Stmt
	makeInactiveStmt	  *sql.Stmt
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

	getAllUserStmt, err := db.Prepare("SELECT email, is_active FROM Users")
	if err != nil {
		return nil, err
	}

	makeActiveStmt, err := db.Prepare("UPDATE Users SET is_active = 1 WHERE email = ?")
	if err != nil {
		return nil, err
	}

	makeInactiveStmt, err := db.Prepare("UPDATE Users SET is_active = 0 WHERE email = ?")
	if err != nil {
		return nil, err
	}

	return &MySQLUserRepository{
		db:                    db,
		insertStmt:            insertStmt,
		getUserStmt: 		   getUserStmt,
		getAllUserStmt:		   getAllUserStmt,
		makeActiveStmt: 	   makeActiveStmt,
		makeInactiveStmt: 	   makeInactiveStmt,
	}, nil
}

func (repo *MySQLUserRepository) Insert(user *models.User) (*models.User, error) {
	_, err := repo.insertStmt.Exec(user.Email, user.IsActive)
	if err != nil {
		return nil, err
	}
	return user, nil
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

func (repo *MySQLUserRepository) GetAllUser() ([]models.User, error) {
	rows, err := repo.getAllUserStmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Email, &user.IsActive)
		if err != nil {
			fmt.Println("Error rows:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error rows2:", err)
		return nil, err
	}

	return users, nil
}

func (repo *MySQLUserRepository) MakeActive(email string) error {
	_, err := repo.makeActiveStmt.Exec(email)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MySQLUserRepository) MakeInactive(email string) error {
	_, err := repo.makeInactiveStmt.Exec(email)
	if err != nil {
		return err
	}
	return nil
}