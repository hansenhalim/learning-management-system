package entity

type Classroom struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`        // e.g., Room 101
	TeacherID  uint   `json:"teacher_id"`  // Links to Teacher
	StudentIDs []uint `json:"student_ids"` // Links to Students
	Capacity   int    `json:"capacity"`    // Maximum number of students
}
