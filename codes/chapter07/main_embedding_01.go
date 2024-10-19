package main

import "fmt"

type Employee struct {
	Name string
	Id   int
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%d)", e.Name, e.Id)
}

type Manager struct {
	Employee // EMBEDDED
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return m.Reports
}

func main() {
	m := Manager{
		Employee: Employee{"John", 100},
		Reports: []Employee{
			Employee{"Alice", 200},
			Employee{"Bob", 201},
		},
	}

	fmt.Println(m.Description())      // John (100)
	fmt.Println(m.FindNewEmployees()) // [{Alice 200} {Bob 201}]
}
