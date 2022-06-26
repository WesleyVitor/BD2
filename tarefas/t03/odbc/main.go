package main

import (
	"github.com/WesleyVitor/dcl"
	"github.com/WesleyVitor/dml"
)


func main(){
	
	dcl.CreateDB()
	dml.InserirFuncionario()
	dml.InserirProjeto()
	dml.InserirAtividade()
	dml.UpdateResponsavel()
}