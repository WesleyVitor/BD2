package dml

import (
	"fmt"

	"github.com/WesleyVitor/dcl"
)


func InserirFuncionario(){
	db := dcl.Connect()

	insert := fmt.Sprint(`
		INSERT INTO funcionario(
			nome, sexo, dtNasc, salario, codSupervisor, codDepto)
		values ('Maria', 'F', '1981-07-01', 2500.00, null, null);
		
		INSERT INTO funcionario(
			nome, sexo, dtNasc, salario, codSupervisor, codDepto)
		values ('José', 'M', '1981-07-01', 1500.00, null, null);
		
	`)
	_,err := db.Exec(insert)
	dcl.HandlerError(err)

}