package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ResultUser(c *gin.Context, db *gorm.DB) {
	// ac := gin.New()
	var user []User
	db.Find(&user)
	c.JSON(http.StatusOK, user)

}
func Savedata(c *gin.Context, db *gorm.DB) {
	data := User{
		Name: "satriAa",
	}
	err := db.Create(&data).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}
