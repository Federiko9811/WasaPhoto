package database

func (db *appdbimpl) PostPhoto(image []byte, token int64) error {
	_, err := db.c.Exec("INSERT INTO photo (owner, img) VALUES (?, ?)", token, image)
	return err
}

func (db *appdbimpl) DeletePhoto(token int64, photoId int64) error {
	_, err := db.c.Exec("DELETE FROM photo WHERE owner=? AND id=?", token, photoId)
	return err
}

func (db *appdbimpl) GetImage(photoId int64) ([]byte, error) {
	var image []byte
	err := db.c.QueryRow("SELECT img FROM photo WHERE id=?", photoId).Scan(&image)
	return image, err
}

func (db *appdbimpl) GetPhotoOwner(photoId int64) (int64, error) {
	var owner int64
	err := db.c.QueryRow("SELECT owner FROM photo WHERE id=?", photoId).Scan(&owner)
	return owner, err
}

func (db *appdbimpl) CheckPhotoOwner(token int64, photoId int64) (bool, error) {
	var owner int64
	err := db.c.QueryRow("SELECT owner FROM photo WHERE id=?", photoId).Scan(&owner)
	if err != nil {
		return false, err
	}
	return owner == token, nil
}

func (db *appdbimpl) LikePhoto(token int64, photoId int64) error {
	_, err := db.c.Exec("INSERT INTO like (owner, photo) VALUES (?, ?)", token, photoId)
	return err
}

func (db *appdbimpl) UnlikePhoto(token int64, photoId int64) error {
	_, err := db.c.Exec("DELETE FROM like WHERE owner=? AND photo=?", token, photoId)
	return err
}

func (db *appdbimpl) CheckLike(token int64, photoId int64) (bool, error) {
	var count int64
	err := db.c.QueryRow("SELECT count(*) FROM like WHERE owner=? AND photo=?", token, photoId).Scan(&count)
	if err != nil || count == 0 {
		return false, err
	}
	return count == 1, nil
}
