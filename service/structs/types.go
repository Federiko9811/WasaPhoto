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
