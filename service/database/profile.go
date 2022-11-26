package database

import (
	"database/sql"
	"errors"
)

// GetUserToken returns the user token for the given username.
func (db *appdbimpl) GetUserToken(username string) (int64, error) {
	var token int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT token FROM user WHERE username=?", username).Scan(&token)

	if errors.Is(err, sql.ErrNoRows) {
		// If no user has been found, create a new one
		token, err = db.addUser(username)
		if err != nil {
			return 0, err
		}
	}

	return token, nil
}

// addUser adds a new user to the database and returns the user token.
func (db *appdbimpl) addUser(username string) (int64, error) {
	var token int64

	row, err := db.c.Exec("INSERT INTO user (username) VALUES (?) RETURNING token", username)

	if err != nil {
		return 0, err
	}

	token, err = row.LastInsertId()
	return token, err
}

// SetUserName sets the username for the given user token.
func (db *appdbimpl) SetUserName(token int64, username string) error {
	_, err := db.c.Exec("UPDATE user SET username=? WHERE token=?", username, token)
	if err != nil {
		return err
	}
	return nil
}

// CheckToken checks if the token exist.
func (db *appdbimpl) CheckToken(token int64) bool {
	var count int64
	err := db.c.QueryRow("SELECT COUNT(*) FROM user WHERE token=?", token).Scan(&count)
	if err != nil {
		return false
	}
	return count == 1
}
