package dto

type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
	Role      string `json:"role"`
	Password  string `json:"password"`
}
