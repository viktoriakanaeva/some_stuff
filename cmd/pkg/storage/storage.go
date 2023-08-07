package storage

import (
	"fmt"

	"github.com/viktoriakanaeva/cmd/pkg/student"
)

type Storage map[string]*student.Student

func NewStorage() Storage {
	return make(Storage)
}
func (sto Storage) Put(stu *student.Student) {

	sto[stu.Name] = stu

}
func (sto Storage) Get(UserName string) (stu *student.Student) {

	fmt.Println("Студенты из хранилища:")
	for _, stu := range sto {
		fmt.Println(stu)

	}
	return
}
