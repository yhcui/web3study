package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB
var err error

type Employee struct {
	Id         int
	Name       string
	Department string
	Salary     int
}

func main() {
	initDB()
	queryDep()
	queryMaxSalary()

}

func queryMaxSalary() {
	var employees Employee
	err = dbx.Get(&employees, "SELECT id, name, department, salary FROM employees order by salary desc limit 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("============最高工资")
	fmt.Println(employees)
}

func queryDep() {
	var employees []Employee
	err := dbx.Select(&employees, "SELECT id, name, department, salary FROM employees WHERE department=?", "技术部")
	if err != nil {
		log.Fatal("Query failed:", err)
	}

	// 输出结果
	fmt.Println("Query result:")
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %d\n", emp.Id, emp.Name, emp.Department, emp.Salary)
	}
}
func initDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"
	dbx, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	dbx.SetMaxOpenConns(30)
	dbx.SetMaxIdleConns(10)
	return
}
