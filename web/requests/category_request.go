package requests

type AddCategoryRequest struct {
	Category string `validate:"required"`
}
