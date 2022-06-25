package dcl

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Pessoa struct{
	ID int
	Nome string
}




func Model(){


	connStr := "user=postgres dbname=postgres password=postgres host=172.20.0.3 sslmode=disable"
	
	db, err := sql.Open("postgres", connStr)
	
	if err != nil{
		panic("Deu RUIM!")
	}

	rows, err := db.Query(`select * from pessoa;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var p Pessoa
		err := rows.Scan(&p.ID, &p.Nome)
		if err != nil {
			panic("Deu ruim no for")
		}
		fmt.Println(p.ID, p.Nome)
	}

	
	
}