package user_models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
