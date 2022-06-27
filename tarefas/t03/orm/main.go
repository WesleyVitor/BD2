package main

import (
	"github.com/WesleyVitor/connection"
	"github.com/WesleyVitor/models"
)
func main(){
	db :=connection.Connect()
	
	db.AutoMigrate(&models.Departamento{})
	db.AutoMigrate(&models.Funcionario{})
	db.AutoMigrate(&models.Projeto{})
	db.AutoMigrate(&models.Atividade{})
	
	


}