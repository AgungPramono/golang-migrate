package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=100,min=3" json:"name"`
}
