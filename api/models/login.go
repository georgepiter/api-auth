package models

// Login representa um login.
// swagger:response loginResponse
type Login struct {
	LoginName string `json:"userName,omitempty" example:"piter"`
	Email     string `json:"email,omitempty" example:"piter.teste@example.com"`
	Password  string `json:"password,omitempty" example:"123"`
}
