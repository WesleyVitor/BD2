package main

import (
	"time"

	"github.com/WesleyVitor/connection"
	"github.com/WesleyVitor/models"
	"github.com/WesleyVitor/repository"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB){
	
	db.AutoMigrate(&models.Departamento{})
	db.AutoMigrate(&models.Funcionario{})
	db.AutoMigrate(&models.Projeto{})
	db.AutoMigrate(&models.Atividade{})
}

func main(){
	db :=connection.Connect()
	migrate(db)

	projeto := models.Projeto{Nome:"SIGARTE",Descricao:"Projeto de artesões",
	DataInicio:time.Date(2022,8,11,0,0,0,0,time.UTC),DataFim:time.Date(2022,12,11,0,0,0,0,time.UTC)}
	repository.CriarProjeto(&projeto,db)
	atividade := models.Atividade{Descricao:"Fazer o CRUD",CodProjeto:1,
	DataInicio:time.Date(2022,8,11,0,0,0,0,time.UTC),DataFim:time.Date(2022,12,11,0,0,0,0,time.UTC)}
	repository.CriarAtividade(&atividade, db)

	
	

	


}