package main

import ("fmt")

type Employee struct {
	Salary int
	Name   string
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func main() {
	employee := Employee{Salary: 30000, Name: "Bruce"}

	fmt.Println(Bonus(&employee, 15))
	AwardAnnualRaise(&employee)
	fmt.Println(employee.Salary)
}
