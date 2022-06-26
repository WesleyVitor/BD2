package dcl

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const host     = "172.20.0.1"
const user 	   = "postgres"
const password = "postgres"
const dbname   = "go_db"
const port 	   = "5432"


func Connect()*sql.DB{
	ctn := fmt.Sprintf("user=%s port=%s dbname=%s password=%s host=%s sslmode=disable",
					user,port,dbname,password,host)
	
	db, err := sql.Open("postgres", ctn)
	
	if err != nil{
		log.Fatal(err)
	}

	return db
}

// func readingRows(rows *sql.Rows){
// 	for rows.Next() {
// 		var ID string
// 		var Nome string
// 		err := rows.Scan(&ID, &Nome)
// 		if err != nil {
// 			panic("Deu ruim no for")
// 		}
// 		fmt.Println(ID, Nome)
// 	}
// }

func HandlerError(err error){
	if err != nil {
		log.Fatal(err)
	}
}




func CreateDB(){

	db := Connect()
	defer db.Close()
	creatingTableEmployee := fmt.Sprint(`
		CREATE TABLE IF NOT EXISTS funcionario (
			codigo serial,
			nome varchar(50),
			sexo char(1),
			dtNasc date,
			salario decimal(10,2),
			codSupervisor int,
			codDepto int,
			PRIMARY KEY (codigo),
			FOREIGN KEY (codSupervisor) REFERENCES funcionario(codigo) on delete set null on update cascade
		);
	`)
	_,err := db.Exec(creatingTableEmployee)
	HandlerError(err)


	// rows, err := db.Query(`select * from pessoa;`)
	// HandlerError(err)

	//defer rows.Close()

	
}