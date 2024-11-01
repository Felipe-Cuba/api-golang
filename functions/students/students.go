package students

import (
	"api/models"
	"errors"
	"slices"
)

var students = []models.Student{}

func GetStudents()([]models.Student) {
	return students
}

func InsertStudent(student models.Student) {
	students = append(students, student)
}

func GetStudentById(id int)(models.Student, error) {
	for _, student := range students {
		if student.Id == id{
			return student, nil
		}
	}

	err := errors.New("student not found")
	
	return models.Student{}, err
}

func UpdateStudent(id int, student models.Student) (string, error) {
	actualStudent, err := GetStudentById(id)

	index := slices.Index(students, actualStudent)


	if err != nil {
		return "", err
	}

	if actualStudent.Name != student.Name {
		students[index].Name = student.Name
	}

	if actualStudent.Age != student.Age {
		students[index].Age = student.Age
	}

	return "Student updated successfully!", nil
}

func DeleteStudent(id int) (string, error) {
	actualStudent, err := GetStudentById(id)

	if err != nil {
		return "", err
	}

	index := slices.Index(students, actualStudent)

	students = append(students[:index], students[index+1:]...)

	return "Student deleted successfully!", nil
}

