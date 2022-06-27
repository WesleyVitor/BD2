package models

import (
	"time"

	"gorm.io/gorm"
)


type Funcionario struct {
	gorm.Model
	Codigo uint `gorm:"primaryKey;autoIncrement"`
	Nome string `gorm:"size:255"`
	Sexo string `gorm:"size:2"`
	DtNasc time.Time
	Salario float64
	CodSupervisor string
	CodDepto uint 
}

type Departamento struct {
	gorm.Model
	Codigo uint `gorm:"primaryKey;autoIncrement"`
	Sigla string
	Descricao string
	CodGerente []Funcionario `gorm:"foreignKey:CodDepto"`
}

type Projeto struct{
	gorm.Model
	Codigo uint `gorm:"primaryKey;autoIncrement"`
	Nome string `gorm:"size:255"`
	Descricao string `gorm:"size:255"`
	CodResponsavel Funcionario `gorm:"foreignKey:Codigo"`
	Atividades []Atividade `gorm:"foreignKey:CodProjeto"`
	CodDepto string
	DataInicio time.Time
	DataFim time.Time
	
}

type Atividade struct{
	gorm.Model
	Codigo uint `gorm:"primaryKey;autoIncrement"` 
	Descricao string `gorm:"size:255"`
	CodProjeto uint 
	DataInicio time.Time
	DataFim time.Time
}