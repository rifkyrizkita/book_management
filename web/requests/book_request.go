package requests

// import "time"

type AddNewBookBody struct {
	Title       string `validate:"required"`
	Author      string `validate:"required"`
	Synopsis    string `validate:"required"`
	ISBN        string `validate:"required,len=13"`
	PublishedAt string `validate:"required"`
	Publisher   string `validate:"required"`
	Stock       uint   `validate:"required,min=0"`
	CategoryID  uint   `validate:"required"`
}
