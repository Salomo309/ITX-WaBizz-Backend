package repositories

import (
	"database/sql"

	"itx-wabizz/models"
)

/*** User Repository Declaration and Implementation***/

// Interface for user repository
type UserRepository interface {
	Insert(*models.User) (*models.User, error)
	GetUser(string) (*models.User, error)
	UpdateDeviceToken(string, string) error
	GetAllUser() ([]models.User, error)
	MakeActive(email string) error
	MakeInactive(email string) error
}

// Implementation of user repository
type userRepo struct {
	db              		*sql.DB
	insertStmt       		*sql.Stmt
	getUserStmt      		*sql.Stmt
	updateDeviceTokenStmt	*sql.Stmt
	getAllUserStmt   		*sql.Stmt
	makeActiveStmt   		*sql.Stmt
	makeInactiveStmt 		*sql.Stmt
}

// Function to create new user repository. Prepare all statement and return the instance.
func NewUserRepository(db *sql.DB) (UserRepository, error) {
	insertStmt, err := db.Prepare("INSERT INTO Users (email, is_active, is_admin, device_token) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	getUserStmt, err := db.Prepare("SELECT email, is_active, is_admin, device_token FROM Users WHERE email = ?")
	if err != nil {
		return nil, err
	}

	updateDeviceTokenStmt, err := db.Prepare("UPDATE Users SET device_token = ? WHERE email = ?")
	if err != nil {
		return nil, err
	}

	getAllUserStmt, err := db.Prepare("SELECT email, is_active, is_admin, device_token FROM Users")
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

	return &userRepo{
		db:               		db,
		insertStmt:       		insertStmt,
		getUserStmt:     	 	getUserStmt,
		updateDeviceTokenStmt: 	updateDeviceTokenStmt,
		getAllUserStmt:   		getAllUserStmt,
		makeActiveStmt:   		makeActiveStmt,
		makeInactiveStmt: 		makeInactiveStmt,
	}, nil
}

/*** User Repository Function Implementation ***/

// Insert new user to database. User should not be admin and their device token set to empty string.
func (repo *userRepo) Insert(user *models.User) (*models.User, error) {
	_, err := repo.insertStmt.Exec(user.Email, user.IsActive, user.IsAdmin, user.DeviceToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Find a user from their email. Return error if user not found.
func (repo *userRepo) GetUser(email string) (*models.User, error) {
	row := repo.getUserStmt.QueryRow(email)

	var user models.User
	err := row.Scan(&user.Email, &user.IsActive, &user.IsAdmin, &user.DeviceToken)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update spesific user device registration token.
func (repo *userRepo) UpdateDeviceToken(email string, deviceToken string) error {
	_, err := repo.updateDeviceTokenStmt.Exec(deviceToken, email)
	if err != nil {
		return err
	}

	return nil
}

// Get all user in the database
func (repo *userRepo) GetAllUser() ([]models.User, error) {
	rows, err := repo.getAllUserStmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Email, &user.IsActive, &user.IsAdmin, &user.DeviceToken)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Update a user to be active
func (repo *userRepo) MakeActive(email string) error {
	_, err := repo.makeActiveStmt.Exec(email)
	if err != nil {
		return err
	}

	return nil
}

// Update user to be inactive
func (repo *userRepo) MakeInactive(email string) error {
	_, err := repo.makeInactiveStmt.Exec(email)
	if err != nil {
		return err
	}
	
	return nil
}