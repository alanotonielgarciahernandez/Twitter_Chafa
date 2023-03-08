package models

// LoginResponse, contains JWT
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
