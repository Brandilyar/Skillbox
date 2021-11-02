package student

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(name string, age int, grade int) *Student {
	return &Student{name, age, grade}
}

func (s *Student) ChengeName(newName string){
	s.Name = newName
}