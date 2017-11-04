package demo

import "fmt"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "log"
import "os"

func buildDB(){
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "create table foo(id integer not null primary key, name text);"
	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Println("%q:%s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO foo(id,name) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i:=0;i<100;i++{
		_,err := stmt.Exec(i,fmt.Sprintf("haha%d",i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
 }