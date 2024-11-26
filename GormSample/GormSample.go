package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:12230500o90P@127.0.0.1:1433?database=Godb"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm
	//Server=localhost;Database=Gorm;User Id=SA;Password=12230500o90P;TrustServerCertificate=True
	println(db.Name())
	// Migrate the schema
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1) // find product with integer primary key
	println(product.Code)
}
