package main

import (
	"fmt"
)

type employee struct {
	name   string
	sex    string
	age    int
	salary int
}

func newEmployee(name string, sex string, age int, salary int) employee {
	return employee{
		name:   name,
		age:    age,
		sex:    sex,
		salary: salary,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Person: %s\nAge: %d\nSalary: %d\n", e.name, e.age, e.salary)
}

func main() {
	employees := map[string]employee{
		"1": newEmployee("Milan", "M", 35, 300000),
		"2": newEmployee("Nalim", "W", 12, 102040),
		"3": newEmployee("Yan", "M", 24, 20500),
	}

	// Печать информации о всех сотрудниках
	for id, emp := range employees {
		fmt.Printf("ID: %s\n%s\n", id, emp.getInfo())
	}

	// Удаление сотрудника
	var key string
	fmt.Print("Enter the ID of the employee to delete: ")
	_, err := fmt.Scanf("%s\n", &key)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if _, exists := employees[key]; exists {
		delete(employees, key)
		fmt.Printf("Employee with ID %s has been deleted.\n", key)
	} else {
		fmt.Printf("Employee with ID %s does not exist.\n", key)
	}

	// Печать оставшихся сотрудников
	fmt.Println("\nRemaining employees:")
	for id, emp := range employees {
		fmt.Printf("ID: %s\n%s\n", id, emp.getInfo())
	}
}
