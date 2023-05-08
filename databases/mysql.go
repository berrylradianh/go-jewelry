package databases

import (
	"fmt"
	"os"

	e "github.com/berrylradianh/go-jewelry/modules/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
}

func InitDB() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		e.ProductCategory{},
		e.ProductDescription{},
		e.ProductMaterial{},
		e.Product{},
		e.Cart{},
		e.User{},
		e.UserDetail{},
		e.Role{},
		e.Payment{},
		e.Transaction{},
		e.TransactionDetail{},
	)
}
