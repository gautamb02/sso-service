package user

type UserDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Verified  bool   `json:"verified"`
}
