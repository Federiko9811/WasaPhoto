package database

import (
	"WasaPhoto/service/structs"
	"database/sql"
	"errors"
)

// MAIN FUNCTION

// GetUserToken returns the user token for the given username.
func (db *appdbimpl) GetUserToken(username string) (int64, error) {

	token, err := db.GetUserTokenOnly(username)

	if errors.Is(err, sql.ErrNoRows) {
		// If no user has been found, create a new one
		token, err = db.addUser(username)
		if err != nil {
			return 0, err
		}
	}

	return token, nil
}

// GetUserProfile returns the user profile for the given user token.
func (db *appdbimpl) GetUserProfile(us string) (structs.UserProfile, error) {
	var profile structs.UserProfile

	// Get the user token from the database
	var token int64
	err := db.c.QueryRow("SELECT token FROM user WHERE username=?", us).Scan(&token)
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

	profile.NumberOfFollowers, err = db.getNumberFollowers(profile.Token)
	if err != nil {
		return profile, err
	}

	profile.NumberOfFollowing, err = db.getNumberFollowing(profile.Token)
	if err != nil {
		return profile, err
	}

	return profile, nil

}

// SetUserName sets the username for the given user token.
func (db *appdbimpl) SetUserName(token int64, username string) error {
	_, err := db.c.Exec("UPDATE user SET username=? WHERE token=?", username, token)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetUsersList(str string) ([]string, error) {
	var users []string
	// Get all the photos of the user
	rows, err := db.c.Query("SELECT username FROM user WHERE username LIKE ?", "%"+str+"%")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			return nil, err
		}
		users = append(users, username)
	}

	if rows.Err() != nil {
		return users, rows.Err()
	}

	defer rows.Close()

	return users, nil
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

func (db *appdbimpl) GetNumberOfLikes(photoId int64) (int64, error) {
	var count int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT count(*) FROM like WHERE photo=?", photoId).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (db *appdbimpl) GetNumberOfComments(photoId int64) (int64, error) {
	var count int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT count(*) FROM comment WHERE photo=?", photoId).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (db *appdbimpl) getNumberFollowers(token int64) (int64, error) {
	var count int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT count(*) FROM follow WHERE followed=?", token).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (db *appdbimpl) getNumberFollowing(token int64) (int64, error) {
	var count int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT count(*) FROM follow WHERE following=?", token).Scan(&count)
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
		photo.NumberOfLikes, err = db.GetNumberOfLikes(photo.Id)
		if err != nil {
			return nil, err
		}

		// Get the number of comments for the photo
		photo.NumberOfComments, err = db.GetNumberOfComments(photo.Id)
		if err != nil {
			return nil, err
		}

		photo.IsLiked = false

		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return photos, rows.Err()
	}

	defer rows.Close()

	return photos, nil
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

func (db *appdbimpl) GetUserTokenOnly(username string) (int64, error) {
	var token int64
	err := db.c.QueryRow("SELECT token FROM user WHERE username=?", username).Scan(&token)
	if err != nil {
		return -1, err
	}
	return token, nil
}
