package dto

// for auth like register and login
type AuthUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token   string            `json:"token"`
	Message map[string]string `json:"message,omitempty"`
}
