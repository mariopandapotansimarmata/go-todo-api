package web

type TodoCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
