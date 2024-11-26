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

type CreateUserDto struct{
	UserName string
}

func setupRouter() *gin.Engine {

	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:12230500o90P@127.0.0.1:1433?database=GormTest"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	println(db.Name())

	db.AutoMigrate(&User{})

	r := gin.Default()

	r.POST("/adduser",func(c *gin.Context){
		var user CreateUserDto
		
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&User{UserName: user.UserName})
		c.JSON(http.StatusOK, gin.H{"name": user.UserName})
	})

	r.GET("/getuser", func(c *gin.Context) {
		var user User
		db.Find(&user,user.UserName=="sina")
		c.JSON(http.StatusOK, gin.H{"name":user.UserName,"Id":user.ID})
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
