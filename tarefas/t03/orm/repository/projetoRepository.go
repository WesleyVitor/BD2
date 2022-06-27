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
