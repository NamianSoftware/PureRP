package api

type Character struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserId    int    `json:"user_id"`
}
