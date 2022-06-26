package dml

import (
	"fmt"

	"github.com/WesleyVitor/dcl"
)


func UpdateResponsavel(){
	db := dcl.Connect()

	insert := fmt.Sprint(`
		UPDATE projeto
		SET codResponsavel = 3
		WHERE nome ='APF';
	`)
	_,err := db.Exec(insert)
	dcl.HandlerError(err)

}