package repository

import (
	"github.com/WesleyVitor/models"
	"gorm.io/gorm"
)

func CriarAtividade(a *models.Atividade, db *gorm.DB) {
	
	db.Create(&a)
	
}
