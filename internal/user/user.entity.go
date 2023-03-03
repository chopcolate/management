package user

type User struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Role     []string `json:"role"`
}
