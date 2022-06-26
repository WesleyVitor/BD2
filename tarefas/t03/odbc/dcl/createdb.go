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

func handleEmployee(db *sql.DB){
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
}

func handleDept(db *sql.DB){
	creatingTableDept := fmt.Sprint(`
		CREATE TABLE IF NOT EXISTS departamento (
			codigo serial,
			sigla varchar(10),
			descricao varchar(50),
			codGerente int,
			PRIMARY KEY (codigo),
			UNIQUE(sigla),
			FOREIGN KEY (codGerente) REFERENCES funcionario(codigo) on delete set null on update cascade
		);
	`)
	_,err := db.Exec(creatingTableDept)
	HandlerError(err)
}

func handleProject(db *sql.DB){
	creatingTableProject := fmt.Sprint(`
		CREATE TABLE IF NOT EXISTS projeto (
			codigo serial,
			nome varchar(50),
			descricao varchar(250),
			codResponsavel int,
			codDepto int,
			dataInicio date, 
			dataFim date,
			PRIMARY KEY (codigo),
			UNIQUE(nome),
			FOREIGN KEY (codResponsavel) REFERENCES funcionario(codigo) on delete set null on update cascade,
			FOREIGN KEY (codDepto) REFERENCES departamento(codigo) on delete set null on update cascade
		);
	`)
	_,err := db.Exec(creatingTableProject)
	HandlerError(err)
}

func handleActivity(db *sql.DB){
	creatingTableActivity := fmt.Sprint(`
		CREATE TABLE IF NOT EXISTS atividade (
			codigo serial,
			descricao varchar(250),
			codProjeto int,
			dataInicio date, 
			dataFim date,
			PRIMARY KEY (codigo),
			FOREIGN KEY (codProjeto) REFERENCES projeto(codigo) on delete set null on update cascade

		);
	`)
	_,err := db.Exec(creatingTableActivity)
	HandlerError(err)
}

func changeEmployeeADDForeignKEY(db *sql.DB){
	command := fmt.Sprint(`
		alter table funcionario ADD CONSTRAINT funcDeptoFK FOREIGN KEY (codDepto) 
		REFERENCES departamento(codigo) on delete set null on update cascade;
	`)
	_,err := db.Exec(command)
	HandlerError(err)
}





func CreateDB(){

	db := Connect()
	defer db.Close()
	
	handleEmployee(db)
	handleDept(db)
	handleProject(db)
	handleActivity(db)
	changeEmployeeADDForeignKEY(db)

	// rows, err := db.Query(`select * from pessoa;`)
	// HandlerError(err)

	//defer rows.Close()

	
}