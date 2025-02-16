package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	// create a database object which can be used
	// to connect with database.
	db, err := sql.Open("mysql", "root:1234@tcp(0.0.0.0:3306)/customer")
	
if err !=nil{
	panic(err)
}

err = db.Ping()

if err!=nil{
	panic(err)
}

result,err:=db.Query("select * from users")

if err !=nil{
	panic(err)	
}

for result.Next(){

	
	var id int 
	var name string
	var age  int

	result.Scan(&id,&name,&age)

	if err !=nil{
		panic(err)
	}


	

	fmt.Printf("id: %d ,name: %s,age: %d \n",id,name,age )
	
	
}
fmt.Print(" data is sucessfully displayed on terminal")

}
