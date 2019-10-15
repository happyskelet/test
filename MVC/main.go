package main

import (
	con "./controllers"
	"github.com/gin-gonic/gin"
)


var router *gin.Engine

func main(){
	router = gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.GET("/",con.GetHome);
	router.GET("/about",con.GetAbout);
	router.GET("/product/:id",con.GetProduct);
	router.GET("/product",con.GetProducts);
	router.POST("/addProduct",con.AddProduct);
	router.POST("/updateProduct",con.UpdateProduct);
	router.GET("/deleteProduct/:id",con.DeleteProduct);
	router.POST("/upload",con.Upload);
	router.POST("/json",con.JSON);
	router.Run()
}