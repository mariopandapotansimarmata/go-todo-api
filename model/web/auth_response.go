package web

type AuthResponse struct {
	Token   string `validate:"required" json:"token"`
	Expires string `validate:"required" json:"expires"`
}
