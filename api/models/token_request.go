package models

// TokenValidRequest representa um TokenValidRequest.
// swagger:response TokenValidResponse
type TokenValidRequest struct {
	Token string `json:"token,omitempty"`
}
