package structs

type Username struct {
	Username string `json:"name"`
}

type Token struct {
	Identifier int64 `json:"identifier"`
}

type Message struct {
	Message string `json:"message"`
}

type Comment struct {
	Comment string `json:"comment"`
}

type Photo struct {
	Id               int64  `json:"id"`
	Owner            int64  `json:"owner"`
	CreatedAt        string `json:"createdAt"`
	NumberOfLikes    int64  `json:"numberOfLikes"`
	NumberOfComments int64  `json:"numberOfComments"`
}

type UserProfile struct {
	Token             int64   `json:"token"`
	Username          string  `json:"username"`
	Photos            []Photo `json:"photos"`
	NumberOfPhotos    int64   `json:"numberOfPhotos"`
	NumberOfFollowers int64   `json:"numberOfFollowers"`
	NumberOfFollowing int64   `json:"numberOfFollowing"`
}
