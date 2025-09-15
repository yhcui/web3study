package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbt *sql.DB
var err error

func main() {
	dbt, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	b := transaction(500)
	log.Println("转账", b)
}

func transaction(amount int) bool {
	var tamount int
	begin, err3 := dbt.Begin()
	if err3 != nil {
		log.Fatal(err3)
	}
	begin.QueryRow("select balance from accounts where id = ?", "A").Scan(&tamount)
	if tamount < amount {
		log.Printf("钱不够")
		begin.Rollback()
		return false
	}
	prepare, err3 := begin.Prepare("update accounts set balance = balance - ?  where id = ? and balance > ?")
	if err3 != nil {
		log.Fatal(err3)
		begin.Rollback()
		return false
	}
	exec, err3 := prepare.Exec(amount, "A", amount)
	affected, err3 := exec.RowsAffected()

	if affected > 0 {
		result, err3 := begin.Exec("INSERT INTO transactions (`from_account_id`, `to_account_id`, `amount`) VALUES (?, ?, ?)", "A", "B", amount)
		if err3 != nil {
			begin.Rollback()
			log.Fatal("err3")
			log.Fatal(err3)
			return false
		}
		rowsAffected, err3 := result.RowsAffected()
		if err3 != nil {
			begin.Rollback()
			log.Fatal(err3)
			return false
		}
		if rowsAffected > 0 {
			begin.Commit()
			return true
		} else {
			begin.Rollback()

		}
	}
	return false
}
