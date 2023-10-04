package users

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" pg:",unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}