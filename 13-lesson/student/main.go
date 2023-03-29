package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BookData interface {
	GetStudent() []Student
	GetStudentByID(id int) (*Student, error)
	AddStudent(student Student) error
	RemoveStudent(id int)
	UpdateStudent(id int)
}

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Section string `json:"section"`
	Course  string `json:"course"`
	Subject []Sub  `json:"subjects"`
}

var id, subjectId, result, marks int
var studentName, subjectName, section, course string

type Sub struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Marks int    `json:"marks"`
}

type AllStudent struct {
	Students []Student
}

func (s *AllStudent) GetStudent() []Student {
	fmt.Println("Read")
	// for _, v := range s.Students {
	// 	fmt.Println(v)
	// }
	return s.Students
}

// func (s *AllStudent) AddStudent(student Student) error {
// 	s.Students = append(s.Students, student)
// 	return nil
// }

func (s *AllStudent) AddStudent(student Student) int {

	//student.ID = len(studentName)+1

	fmt.Println("Read onlyr rr", student.Subject)
	for _, k := range student.Subject {
		result = k.Marks
		fmt.Println("fffjfjffjjfjfjfjjfjfjfjjf", result)
		fmt.Println("V", k.Marks)
	}
	s.Students = append(s.Students, student)
	return result
}

func (lib *AllStudent) GetStudentByID(student Student) (*Student, error) {
	for _, b := range lib.Students {
		if b.ID == student.ID {
			return &student, nil
		}
	}
	return &student, nil
}

func main() {
	data, err := ioutil.ReadFile("info.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var subject AllStudent
	err = json.Unmarshal(data, &subject)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	student := AllStudent{Students: subject.Students}
	//fmt.Println(student)
	student.GetStudent()

	newStudent := Student{
		ID:      id,
		Name:    "studentName",
		Grade:   student.AddStudent(student.Students[result]),
		Section: section,
		Course:  course,
		Subject: []Sub{
			{
				ID:    1,
				Name:  "Maths",
				Marks: 80, // mark+mark+mark = grade
			},
			{
				ID:    2,
				Name:  "Science",
				Marks: 70, // mark+mark+mark = grade
			},
			{
				ID:    3,
				Name:  "English",
				Marks: 90, // mark+mark+mark = grade
			},
		},
	}
	student.AddStudent(newStudent)

	data, err = json.Marshal(student)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	err = ioutil.WriteFile("info.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	student.GetStudent()
	fmt.Println("new student", student.GetStudent())

}
