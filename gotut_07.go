package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAppPage struct {
	Title string
	News  string
}

type Employee struct {
	EmpId         string
	EmpName       string
	EmpDepartment string
}

type EmployeePage struct {
	Title        string
	Employees    map[string]Employee
	ArrEmployees []Employee
}

func employeeDetails(w http.ResponseWriter, r *http.Request) {
	employees := []Employee{
		{
			EmpId:         "EMP001",
			EmpName:       "Rishikesh",
			EmpDepartment: "Product Development",
		},
		{
			EmpId:         "EMP002",
			EmpName:       "Nikhil",
			EmpDepartment: "Project Development",
		},
		{
			EmpId:         "EMP003",
			EmpName:       "Akshay",
			EmpDepartment: "Sales Development",
		},
	}

	employee_map := make(map[string]Employee)

	for idx, _ := range employees {
		employee_map[employees[idx].EmpId] = Employee{employees[idx].EmpId, employees[idx].EmpName, employees[idx].EmpDepartment}
	}

	fmt.Println(employees)
	fmt.Println(employee_map)
	p := EmployeePage{Title: "Employee App Title", Employees: employee_map, ArrEmployees: employees}
	t, _ := template.ParseFiles("employeehtml.html")
	fmt.Println(t.Execute(w, p))
}

func aggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAppPage{Title: "News App Title", News: "This is a news"}
	t, _ := template.ParseFiles("agghtml.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey, HTML from go tut</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg", aggHandler)
	http.HandleFunc("/employees", employeeDetails)
	http.ListenAndServe(":8000", nil)
}
