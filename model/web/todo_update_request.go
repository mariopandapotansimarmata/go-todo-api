package web

type TodoUpdateRequest struct {
	Id   int    `validate:"required" `
	Name string `validate:"required" json:"name"`
}
