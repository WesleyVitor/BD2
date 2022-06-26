package dml

import (
	"fmt"

	"github.com/WesleyVitor/dcl"
)


func InserirProjeto(){
	db := dcl.Connect()

	insert := fmt.Sprint(`
		insert into projeto(
			nome, descricao, codDepto, codResponsavel, dataInicio, dataFim)
		values ('APF', 'Analisador de Ponto de Função', null, 1, '2018-02-26', '2019-06-30');	
	
	`)
	_,err := db.Exec(insert)
	dcl.HandlerError(err)

}