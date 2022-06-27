package dcl

import "database/sql"


type Funcionario struct {
	Codigo int
	Nome string
	Sexo string
	DtNasc string
	Salario float64
	CodSupervisor sql.NullInt64
	CodDepto sql.NullInt64
}

type Departamento struct {
	Codigo int 
	Sigla string
	Descricao string
	CodGerente sql.NullInt64
}

type Projeto struct {
	Codigo int
	Nome string
	Descricao string
	CodResponsavel sql.NullInt64
	CodDepto sql.NullInt64
	DataInicio string
	DataFim string
	
}

type Atividade struct{
	Codigo int 
	Descricao string
	CodProjeto sql.NullInt64
	DataInicio string
	DataFim string
}