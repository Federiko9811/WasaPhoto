package database

import (
	"WasaPhoto/service/structs"
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	Ping() error
	GetUserTokenOnly(username string) (int64, error)
	CheckPhotoOwner(token int64, photoId int64) (bool, error)
	GetPhotoOwner(photoId int64) (int64, error)
	CheckLike(token int64, photoId int64) (bool, error)
	GetNumberOfLikes(photoId int64) (int64, error)
	GetNumberOfComments(photoId int64) (int64, error)
	CheckPhotoExistence(photoId int64) (bool, error)

	GetUserToken(username string) (int64, error)
	SetUserName(token int64, username string) error
	CheckToken(token int64) bool
	GetUserProfile(username string, requestUser int64) (structs.UserProfile, error)
	GetUsersList(username string) ([]string, error)

	AddFollow(following int64, follow string) error
	RemoveFollow(following int64, follow string) error
	AddBan(banning int64, ban string) error
	RemoveBan(banning int64, ban string) error
	CheckFollow(u1 int64, u2 int64) (bool, error)
	CheckBan(u1 int64, u2 int64) (bool, error)

	PostPhoto(image []byte, token int64) error
	DeletePhoto(token int64, photoId int64) error
	GetImage(photoId int64) ([]byte, error)
	LikePhoto(token int64, photoId int64) error
	UnlikePhoto(token int64, photoId int64) error
	CommentPhoto(token int64, photoId int64, content string) (int64, error)
	GetPhotoComments(photoId int64) ([]structs.FullDataComment, error)
	GetCommentOwner(commentId int64) (int64, error)
	GetComment(commentId int64) (structs.FullDataComment, error)
	DeleteComment(commentId int64) error
	GetMyStream(token int64) ([]structs.Photo, error)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
