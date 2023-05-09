package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqlconfig := mysql.Config{}
	mysqlconfig.Addr = "localhost:3306"
	mysqlconfig.Net = "tcp"
	mysqlconfig.User = "root"
	mysqlconfig.Passwd = "cbt123"

	db, err := sql.Open("mysql", mysqlconfig.FormatDSN())
	fmt.Println(db, err)

	result, err := db.Exec("CREATE DATABASE IF NOT EXISTS cbtnuggets;")
	fmt.Println(result, err)

	tableresult, err := db.Exec("CREATE TABLE IF NOT EXISTS cbtnuggets.people (FirstName varchar(30), LastName varchar(30))")
	fmt.Println(tableresult, err)

	insertresult, err := db.Exec("INSERT INTO cbtnuggets.people (FirstName, LastName) VALUES ('Trevor', 'Sullivan'), ('John', 'McGovern')")
	fmt.Println(insertresult, err)

	select_query := "SELECT LastName,FirstName FROM cbtnuggets.people;"
	selectresult, err := db.Query(select_query)
	fmt.Println(selectresult, err)

	allpeople := []Person{}
	for selectresult.Next() {
		newPerson := Person{}
		selectresult.Scan(&newPerson.LastName, &newPerson.FirstName)
		allpeople = append(allpeople, newPerson)
	}

	fmt.Println(allpeople)
}

type Person struct {
	FirstName string
	LastName  string
}
