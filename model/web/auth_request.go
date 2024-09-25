package web

type AuthRequest struct {
	UserName string `validate:"required" json:"userName"`
	Password string `validate:"required" json:"password"`
}
