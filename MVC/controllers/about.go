package controllers

import "github.com/gin-gonic/gin"

func GetAbout(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"uuhhuhuhuhu",
		"status":200,	
	})
}