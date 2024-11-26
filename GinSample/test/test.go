package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
}

func setupRouter() *gin.Engine {

	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:12230500o90P@127.0.0.1:1433?database=GormTest"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm
	//Server=localhost;Database=Gorm;User Id=SA;Password=12230500o90P;TrustServerCertificate=True
	println(db.Name())
	// Migrate the schema
	db.AutoMigrate(&User{})

	r := gin.Default()

	r.GET("/adduser", func(c *gin.Context) {
		var name string
		c.ShouldBind(&name)
		db.Create(&User{UserName: "name"})
		c.String(http.StatusOK, "pong")
	})

	r.GET("/getdetails", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Ip": c.ClientIP(), "Name": "kami", "Url": c.Request.URL.Path})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run("192.168.1.20:8080")
}
