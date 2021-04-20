package entity

// Repository operate database by entities
type Repository interface {
	Save() error
}
