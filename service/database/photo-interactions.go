package database

import (
	"WasaPhoto/service/structs"
)

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

func (db *appdbimpl) CommentPhoto(token int64, photoId int64, content string) error {
	_, err := db.c.Exec("INSERT INTO comment (owner, content, photo) VALUES (?, ?, ?)", token, content, photoId)
	return err
}

func (db *appdbimpl) GetPhotoComments(photoId int64) ([]structs.FullDataComment, error) {
	rows, err := db.c.Query("SELECT * FROM comment WHERE photo=?", photoId)
	if err != nil {
		return nil, err
	}

	var comments []structs.FullDataComment
	for rows.Next() {
		var comment structs.FullDataComment
		err = rows.Scan(&comment.Id, &comment.Content, &comment.CreatedAt, &comment.Owner, &comment.Photo)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return comments, rows.Err()
	}

	defer rows.Close()

	return comments, nil
}

func (db *appdbimpl) GetCommentOwner(commentId int64) (int64, error) {
	var owner int64
	err := db.c.QueryRow("SELECT owner FROM comment WHERE id=?", commentId).Scan(&owner)
	return owner, err
}

func (db *appdbimpl) GetComment(commentId int64) (structs.FullDataComment, error) {
	var comment structs.FullDataComment
	err := db.c.QueryRow("SELECT * FROM comment WHERE id=?", commentId).Scan(&comment.Id, &comment.Content, &comment.CreatedAt, &comment.Owner, &comment.Photo)
	return comment, err
}

func (db *appdbimpl) DeleteComment(commentId int64) error {
	_, err := db.c.Exec("DELETE FROM comment WHERE id=?", commentId)
	return err
}

func (db *appdbimpl) GetMyStream(token int64) ([]structs.Photo, error) {
	rows, err := db.c.Query("select id, owner, created_at from photo where owner not in (select banning from ban where banned=?) and owner != ? and owner in (select followed from follow where following=?) order by created_at desc", token, token, token)
	if err != nil {
		return nil, err
	}

	var photos []structs.Photo
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

		// Check if the photo is liked by the request user
		photo.IsLiked, err = db.CheckLike(token, photo.Id)
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return photos, rows.Err()
	}

	defer rows.Close()

	return photos, nil
}

func (db *appdbimpl) CheckPhotoExistence(photoId int64) (bool, error) {
	var count int64
	err := db.c.QueryRow("SELECT count(*) FROM photo WHERE id=?", photoId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
