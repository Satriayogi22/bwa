package main

import (
	"fmt"

	"bwa/cmd"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	Routing(router, db)
	router.Run(":8111")
}
func Routing(router *gin.Engine, db *gorm.DB) {
	router.GET("/handler", func(c *gin.Context) { cmd.ResultUser(c, db) })
	router.GET("/save", func(c *gin.Context) { cmd.Savedata(c, db) })
}
