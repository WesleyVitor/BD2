package connection

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const host     = "172.20.0.1"
const user 	   = "postgres"
const password = "postgres"
const dbname   = "go_orm"
const port 	   = "5432"

func HandleError(err error){
	if err != nil{
		log.Fatal(err)
	}
}


func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",host,user,password,dbname,port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	HandleError(err)

	return db
}