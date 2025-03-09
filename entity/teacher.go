package entity

type Teacher struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Subject   string `json:"subject"` // e.g., Math, Science
}
