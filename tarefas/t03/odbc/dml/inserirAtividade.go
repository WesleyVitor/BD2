package dml

import (
	"fmt"

	"github.com/WesleyVitor/dcl"
)


func InserirAtividade(){
	db := dcl.Connect()

	insert := fmt.Sprint(`
		insert into atividade(
			descricao, codProjeto, dataInicio, dataFim)
		values ('APF - Atividade 1', 1, '2018-02-26', '2018-06-30');
	`)
	_,err := db.Exec(insert)
	dcl.HandlerError(err)

}