package users

//LoginRequest struct
type LoginRequest struct {
	Email    string `json:"email" sql:"not null;unique"`
	Password string `json:"password"`
}
