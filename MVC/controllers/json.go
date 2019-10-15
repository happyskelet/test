package controllers

import (
	"fmt";
	"github.com/gin-gonic/gin";
	s "../structures"
);

func JSON(c * gin.Context){
	var json s.Json;
	c.ShouldBindJSON(&json);
	fmt.Println(json);
}