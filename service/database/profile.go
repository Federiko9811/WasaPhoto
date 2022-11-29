package database

import (
	"WasaPhoto/service/structs"
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

func (db *appdbimpl) getUserData(id int64) (int64, string, error) {
	var username string
	var token int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT token, username FROM user WHERE token=?", id).Scan(&token, &username)
	if err != nil {
		return -1, "", err
	}

	return token, username, nil
}

func (db *appdbimpl) getNumberOfLikes(photoId int64) (int64, error) {
	var count int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT count(*) FROM like WHERE photo=?", photoId).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (db *appdbimpl) getNumberOfComments(photoId int64) (int64, error) {
	var count int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT count(*) FROM comment WHERE photo=?", photoId).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (db *appdbimpl) getListOfPhotos(token int64) ([]structs.Photo, error) {
	var photos []structs.Photo

	// Get all the photos of the user
	rows, err := db.c.Query("SELECT id, owner, created_at FROM photo WHERE owner=?", token)
	if err != nil {
		return photos, err
	}

	for rows.Next() {
		var photo structs.Photo

		err = rows.Scan(&photo.Id, &photo.Owner, &photo.CreatedAt)
		if err != nil {
			return nil, err
		}

		// Get the number of likes for the photo
		photo.NumberOfLikes, err = db.getNumberOfLikes(photo.Id)
		if err != nil {
			return nil, err
		}

		// Get the number of comments for the photo
		photo.NumberOfLikes, err = db.getNumberOfComments(photo.Id)
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	return photos, nil

}

// TODO Fixare la query
// GetUserProfile returns the user profile for the given user token.
func (db *appdbimpl) GetUserProfile(us string) (structs.UserProfile, error) {
	var profile structs.UserProfile

	// Get the user token from the database
	token, err := db.GetUserToken(us)
	if err != nil {
		return profile, err
	}

	// Get the user profile from the database
	actualToken, username, err := db.getUserData(token)
	if err != nil {
		return profile, err
	}

	profile.Token = actualToken
	profile.Username = username

	profile.Photos, err = db.getListOfPhotos(profile.Token)
	if err != nil {
		return profile, err
	}

	profile.NumberOfPhotos = int64(len(profile.Photos))

	return profile, nil

}
