package validation

type CreateBook struct {
	Name     string `json:"name" validate:"required"`
	AuthorID string `json:"author_id" validate:"required"`
}

type UpdateBook struct {
	Name string `json:"name" validate:"required"`
}
