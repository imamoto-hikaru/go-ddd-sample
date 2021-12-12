package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ApiTransactionManager(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			db, _ := c.Get("db")
			if db != nil {
				db.(*gorm.DB).Rollback()
			}
			c.JSON(200, gin.H{
				"status": "ng",
				"errors": []string{r.(string)},
			})

		}
	}()

	dsn := "host=db user=test password=password dbname=test01 port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		c.AbortWithError(500, err)
	}

	tx := db.Begin()
	c.Set("db", tx)
	c.Next()

	if len(c.Errors) > 0 {
		tx.Rollback()
		errorList := c.Errors
		c.JSON(200, gin.H{
			"status": "ng",
			"errors": errorList.Errors(),
		})
	}

	tx.Commit()
}
