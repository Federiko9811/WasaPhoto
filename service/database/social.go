package database

func (db *appdbimpl) AddFollow(following int64, follow string) error {

	var followed int64

	//TODO USA la query in global.go

	// Get the user token from the database
	err := db.c.QueryRow("SELECT token FROM user WHERE username=?", follow).Scan(&followed)
	if err != nil {
		return err
	}

	// Insert the follow into the database
	_, err = db.c.Exec("INSERT INTO follow (following, followed) VALUES (?, ?)", following, followed)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveFollow(following int64, unfollow string) error {

	var followed int64

	//TODO USA la query in global.go

	// Get the user token from the database
	err := db.c.QueryRow("SELECT token FROM user WHERE username=?", unfollow).Scan(&followed)
	if err != nil {
		return err
	}

	// Delete the follow from the database
	_, err = db.c.Exec("DELETE FROM follow WHERE following=? AND followed=?", following, followed)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) AddBan(banning int64, ban string) error {

	var banned int64

	//TODO USA la query in global.go

	// Get the user token from the database
	err := db.c.QueryRow("SELECT token FROM user WHERE username=?", ban).Scan(&banned)
	if err != nil {
		return err
	}

	// Insert the ban into the database
	_, err = db.c.Exec("INSERT INTO ban (banning, banned) VALUES (?, ?)", banning, banned)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) RemoveBan(banning int64, ban string) error {

	var banned int64

	//TODO USA la query in global.go

	// Get the user token from the database
	err := db.c.QueryRow("SELECT token FROM user WHERE username=?", ban).Scan(&banned)
	if err != nil {
		return err
	}

	//Remove the ban from the database
	_, err = db.c.Exec("DELETE FROM ban WHERE banning=? AND banned=?", banning, banned)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) CheckFollow(u1 int64, u2 int64) (bool, error) {

	count := 0

	err := db.c.QueryRow("SELECT count(*) FROM follow WHERE following=? AND followed=?", u1, u2).Scan(&count)
	if err != nil || count == 0 {
		return false, err
	}
	return true, nil
}

func (db *appdbimpl) CheckBan(u1 int64, u2 int64) (bool, error) {

	count := 0

	err := db.c.QueryRow("SELECT count(*) FROM ban WHERE banning=? AND banned=?", u1, u2).Scan(&count)
	if err != nil || count == 0 {
		return false, err
	}
	return true, nil
}
