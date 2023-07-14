package repo

import "github.com/Tonmoy404/hello-world/service"

type StudentRepo interface {
	service.StudentRepo
}

type studentRepo struct{}

func NewStudentRepo() StudentRepo {
	return &studentRepo{}
}

func (r *studentRepo) GetStudent(id string) *service.Student {
	return &service.Student{
		ID:    "2014751004",
		Name:  "TOnmoy",
		Email: "safdss",
	}
}
