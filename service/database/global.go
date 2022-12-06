package database

func (db *appdbimpl) GetTokenOfUser(username string) (int64, error) {
	var token int64

	// Get the user token from the database
	err := db.c.QueryRow("SELECT token FROM user WHERE username=?", username).Scan(&token)
	if err != nil {
		return 0, err
	}
	return token, nil
}
