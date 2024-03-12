package repositories

import (
	"database/sql"

	"itx-wabizz/models"
)

// Interface for refresh token repository
type RefreshTokenRepository interface {
	Insert(refreshToken *models.RefreshToken) error
	GetRefreshToken(googleID string) (*models.RefreshToken, error)
}

// Implementation of refresh token repository
type MySQLRefreshTokenRepository struct {
	db                  *sql.DB
	insertStmt          *sql.Stmt
	getRefreshTokenStmt *sql.Stmt
}

// Function to create new refresh token repository
func NewMySQLRefreshTokenRepository(db *sql.DB) (*MySQLRefreshTokenRepository, error) {
	insertStmt, err := db.Prepare("INSERT INTO Refresh_Tokens (google_id, refresh_token) VALUES (?, ?) ON DUPLICATE KEY UPDATE refresh_token = VALUES(refresh_token)")
	if err != nil {
		return nil, err
	}

	getRefreshTokenStmt, err := db.Prepare("SELECT google_id, refresh_token FROM Refresh_Tokens WHERE google_id = ?")
	if err != nil {
		return nil, err
	}

	return &MySQLRefreshTokenRepository{
		db:                  db,
		insertStmt:          insertStmt,
		getRefreshTokenStmt: getRefreshTokenStmt,
	}, nil
}

// Function to insert new refresh token into database
func (repo *MySQLRefreshTokenRepository) Insert(refreshToken *models.RefreshToken) error {
	_, err := repo.insertStmt.Exec(refreshToken.Google_ID, refreshToken.Refresh_Token)
	if err != nil {
		return err
	}
	return nil
}

// Function to get refresh token based on the Google ID
func (repo *MySQLRefreshTokenRepository) GetRefreshToken(googleID string) (*models.RefreshToken, error) {
	row := repo.getRefreshTokenStmt.QueryRow(googleID)

	var refreshToken models.RefreshToken
	err := row.Scan(&refreshToken.Google_ID, &refreshToken.Refresh_Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &refreshToken, nil
}
