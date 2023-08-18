package database

import (
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	db := os.Getenv("NAME")
	port := os.Getenv("PORT")
	sslMode := os.Getenv("SSL_MODE")

	stringDeConexao := "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + db +
		" port=" + port +
		" sslmode=" + sslMode

	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
