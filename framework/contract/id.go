package contract

const IDKey = "ggg:id"

type IDService interface {
	NewID() string
}
