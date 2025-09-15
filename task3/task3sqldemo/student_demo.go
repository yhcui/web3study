package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()

	saveSt()
	querySt()
	updateSt()
	deleteSt()
	defer db.Close()
}

func cleanSt(id *int64, name *string) {
	*id = 0
	*name = ""
}

func querySt() {
	var (
		id    int64
		name  string
		age   int
		grade string
	)
	cleanSt(&id, &name)
	rows, err := db.Query("select id, name,age,grade from students where age > ?", 18)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &grade)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age, grade)
	}

}

func deleteSt() {
	stm, err := db.Prepare("delete from students where age < ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stm.Close()
	res, err := stm.Exec(15)
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastId)
	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(affected)
	log.Printf("ID = %d, affected = %d\n", lastId, affected)
}

func updateSt() {
	stm, err := db.Prepare("update students set grade = \"四年级\" where name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stm.Close()
	res, err := stm.Exec("张三")
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastId)
	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(affected)
	log.Printf("ID = %d, affected = %d\n", lastId, affected)
}

func saveSt() {
	stm, err := db.Prepare("INSERT INTO students(name, age, grade) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stm.Close()
	res, err := stm.Exec("张三", 20, "三年级")
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastId)
	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(affected)
	log.Printf("ID = %d, affected = %d\n", lastId, affected)
}
