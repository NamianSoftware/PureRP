package api

type Character struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserCharacter struct {
	Id          int
	UserId      int
	CharacterId int
}
