package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	FIO    string
	Grades []float64
}

// Метод для подсчета среднего балла студента
func (s *Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}

// Ввод данных о студенте
func inputStudent() Student {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите ФИО студента: ")
	scanner.Scan()
	fio := scanner.Text()

	var grades []float64
	fmt.Println("Введите оценки через пробел:")
	scanner.Scan()
	gradesStr := scanner.Text()
	gradesSlice := strings.Fields(gradesStr)

	for _, g := range gradesSlice {
		grade, err := strconv.ParseFloat(g, 64)
		if err != nil {
			fmt.Println("Некорректная оценка, пропускаем.")
			continue
		}
		grades = append(grades, grade)
	}

	return Student{FIO: fio, Grades: grades}
}

// Вывод списка студентов
func printStudents(students map[string]Student) {
	fmt.Println("Список студентов:")
	for key, student := range students {
		fmt.Printf("ID: %s, ФИО: %s, Оценки: %v, Средний балл: %.2f\n", key, student.FIO, student.Grades, student.AverageGrade())
	}
}

func main() {
	students := make(map[string]Student)
	scanner := bufio.NewScanner(os.Stdin)
	idCounter := 1

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1 - Добавить студента")
		fmt.Println("2 - Показать всех студентов")
		fmt.Println("3 - Фильтр по среднему баллу (ниже 4)")
		fmt.Println("4 - Выйти")
		fmt.Print("Выберите действие: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			student := inputStudent()
			id := strconv.Itoa(idCounter)
			students[id] = student
			idCounter++
		case "2":
			printStudents(students)
		case "3":
			fmt.Println("Студенты со средним баллом ниже 4:")
			for id, student := range students {
				if student.AverageGrade() < 4 {
					fmt.Printf("ID: %s, ФИО: %s, Средний балл: %.2f\n", id, student.FIO, student.AverageGrade())
				}
			}
		case "4":
			fmt.Println("Выход.")
			return
		default:
			fmt.Println("Некорректный выбор, попробуйте снова.")
		}
	}
}
