package controllers

import (
	m "../models"
	"github.com/gin-gonic/gin"
	s "../structures"
	"strconv"
)

func GetProduct(c *gin.Context){
	product:=m.GetProduct(c.Param("id"));
	c.JSON(200, product)
}

func GetProducts(c *gin.Context){
	products:=m.GetProducts();
	c.JSON(200,products);
}

func AddProduct(c *gin.Context){
	var test s.Product;
	test.Name = c.PostForm("name");
	test.Description = c.PostForm("description");
	id:=m.AddProduct(test);
	c.JSON(200,gin.H{"id":id,});
}
func UpdateProduct(c *gin.Context){
	var test s.Product;
	test.Id,_ = strconv.Atoi(c.PostForm("id"));
	test.Name = c.PostForm("name");
	test.Description = c.PostForm("description");
	status:=m.UpdateProduct(test);
	c.JSON(200,gin.H{"status":status,});
}
func DeleteProduct(c *gin.Context){
	m.DeleteProduct(c.Param("id"));
	c.JSON(200,gin.H{"product":"deletted"});
}