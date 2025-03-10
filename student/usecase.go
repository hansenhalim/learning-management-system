package student

import "learning-management-system/entity"

//go:generate mockery --name StudentRepository
type StudentRepository interface {
	Find(id uint) (*entity.Student, error)
	Save(student *entity.Student) error
}
