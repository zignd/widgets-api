package viewmodels

const GenericError = "An unexpected error occurred"

type Error struct {
	Message string `json:"message"`
}
