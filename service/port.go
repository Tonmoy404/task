package service

type StudentRepo interface {
	GetStudent(id string) *Student
}

type Service interface {
	GetStudents() []*Student
	GetStudent(ID string) *Student
}
