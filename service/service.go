package service

type service struct {
	studentRepo StudentRepo
}

func NewService(studentRepo StudentRepo) Service {
	return &service{
		studentRepo: studentRepo,
	}
}

func (s *service) GetStudents() []*Student {
	return nil
}

func (s *service) GetStudent(ID string) *Student {
	return s.studentRepo.GetStudent(ID)
}
