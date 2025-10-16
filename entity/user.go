package entity

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
