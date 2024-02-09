package models

// TokenValidResponse representa um TokenValidResponse.
// swagger:response TokenValidResponse
type TokenValidResponse struct {
	IsValid bool   `json:"isValid,omitempty"`
	Token   string `json:"token,omitempty"`
}
