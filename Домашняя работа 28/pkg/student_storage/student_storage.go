package studentstorage

import (
	"fmt"
	"student/pkg/student"
)

type Studentstorage map[string]*student.Student

func NewStudentStorage() Studentstorage {
	return make(map[string]*student.Student)
}

func (s Studentstorage) Put(st *student.Student) {
	s[st.Name] = st
}
func (s Studentstorage) Get(name string) (*student.Student, error) {
	st, ok := s[name]
	if !ok {
		return nil, fmt.Errorf("Студент не найден")
	}
	return st, nil
}
