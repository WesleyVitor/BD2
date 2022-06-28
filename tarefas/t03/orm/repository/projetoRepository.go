package repository

import (
	"github.com/WesleyVitor/models"
	"gorm.io/gorm"
)

func CriarProjeto(p *models.Projeto, db *gorm.DB) {
	result:=db.First(&p, "Nome = ?", p.Nome)
	if result.RowsAffected == 0 {
		db.Create(&p)
	}
	
}

func ListarProjetos(db *gorm.DB) []models.Projeto{
	var projetos []models.Projeto
	db.Table("projetos").Select("codigo, nome, descricao").Scan(&projetos)
	return projetos
}
