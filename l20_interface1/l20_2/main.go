package main

import "fmt"

type Worker interface {
	Work()
}

type Employee struct {
	name string
	age  int
}

func (e Employee) Work() {
	fmt.Println("Now Employee with name:", e.name, "is working!")
}

func Describer(w Worker) {
	fmt.Printf("Interface with type %T and value %v\n", w, w)
}

type Student struct {
	name         string
	courseNumber int
}

func (s Student) Work() {
	fmt.Println("Student with name:", s.name, "is working!")
}

func (s Student) WantToEat() {
	fmt.Println("Student with name:", s.name, "want to eat!")
}

func main() {
	emp := Employee{"Bob", 34}
	var workerEmployee Worker = emp

	workerEmployee.Work()
	Describer(workerEmployee) // type == main.Employee
	fmt.Println()

	std := Student{"Mike", 19}
	var workerStudent Worker = std

	workerStudent.Work()
	Describer(workerStudent) //  type == main.Student
	fmt.Println()

	var workers []Worker = []Worker{workerStudent, workerEmployee}
	for _, worker := range workers {
		Describer(worker)
	}
}
