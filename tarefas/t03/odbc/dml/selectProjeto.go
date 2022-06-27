package dml

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/WesleyVitor/dcl"
)

func showRows(p *dcl.Projeto){
	fmt.Println("Código:", p.Codigo)
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Depto:", p.CodDepto)
	fmt.Println("CodResponsavel:", p.CodResponsavel)
	fmt.Println("Descricao:", p.Descricao)
	fmt.Println("DataInicio:", p.DataInicio)
	fmt.Println("DataFim:", p.DataFim)
	fmt.Println()
}

func readingRows(rows *sql.Rows){
	for rows.Next() {
		var projeto dcl.Projeto
		err := rows.Scan(
			&projeto.Codigo,
			&projeto.Nome,
			&projeto.Descricao,
			&projeto.CodResponsavel,
			&projeto.CodDepto,
			&projeto.DataInicio,
			&projeto.DataFim)
		if err != nil {
			log.Fatal(err)
		}
		showRows(&projeto)
	}
}

func SelectProjeto(){
	db := dcl.Connect()

	query := fmt.Sprint(`
		SELECT * FROM projeto;
	
	`)
	rows, err := db.Query(query)
	dcl.HandlerError(err)
	defer rows.Close()

	readingRows(rows)

}