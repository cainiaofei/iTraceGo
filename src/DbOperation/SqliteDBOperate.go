package DbOperation

import "fmt"
import "database/sql"
import "log"
import 	_ "github.com/mattn/go-sqlite3"


func BuildConnection() {
	dbPath := "/home/zzf/workspace/iTrace4/data/exp/" +
		"Infinispan/rtm/Infinispan-req.db"

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal("...",err)
	}

	sqlStr := "select * from rtm"

	row,err := db.Query(sqlStr)
	fmt.Println("err:",err)
	if err != nil  {
		fmt.Println("dataBase connect failure!!!")
	}else{
		for row.Next() {
			var request string
			var code string
			row.Scan(&request, &code)
			fmt.Println(request,"--->",code)
		}
	}

	defer db.Close()
}

func PrintName(){
	fmt.Println("I am DbOperation")
}

func main(){
	fmt.Println("dd")
}