package entity

type Student struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Grade     string `json:"grade"` // e.g., 10th, 11th, 12th
}
