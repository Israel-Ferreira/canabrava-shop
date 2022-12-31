package config

import (
	"fmt"
	"log"

	"github.com/Israel-Ferreira/canabrava-shop/cadastro-adegas/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s host=%s",
		DbConf.DbUser,
		DbConf.DbPassword,
		DbConf.DbName,
		DbConf.DbPort,
		DbConf.DbHost,
	)

	conn, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalln("Erro ao conectar com o banco: ", err.Error())
	}

	Db = conn

	Db.AutoMigrate(&models.Adega{}, &models.EnderecoAdega{})

}
