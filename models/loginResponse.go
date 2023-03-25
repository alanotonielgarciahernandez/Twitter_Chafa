package models

// LoginResponse, Contains JWT sended when user logs in
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
