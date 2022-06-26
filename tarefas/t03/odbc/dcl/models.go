package dcl

type Funcionario struct {
	Codigo int
	Nome string
	Sexo string
	DtNasc string
	Salario float64
	CodSupervisor int
	CodDepto int
}

type Departamento struct {
	Codigo int 
	Sigla string
	Descricao string
	CodGerente int
}

type Projeto struct {
	Codigo int
	Nome string
	Descricao string
	CodResponsavel int
	CodDepto int
	DataInicio string
	DataFim string
	
}

type Atividade struct{
	Codigo int 
	Descricao string
	CodProjeto int
	DataInicio string
	DataFim string
}