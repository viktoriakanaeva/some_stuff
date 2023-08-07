package main

import (
	"fmt"
	"io"
	"os"

	"github.com/viktoriakanaeva/cmd/pkg/storage"
	"github.com/viktoriakanaeva/cmd/pkg/student"
)

var stud = new(student.Student)

var (
	age   int
	name  string
	grade int
)

func main() {
	sto := storage.NewStorage()
	for {

		fmt.Println("Введите имя студента, возраст и курс")
		_, err := fmt.Fscan(os.Stdin, &name, &age, &grade)
		if err == io.EOF {
			break
		}
		stud = &student.Student{Name: name, Age: age, Grade: grade}

		sto.Put(stud)

	}
	sto.Get(stud.Name)
}
