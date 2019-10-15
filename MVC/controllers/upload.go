package controllers

import ("github.com/gin-gonic/gin";"fmt")

func Upload(c *gin.Context){
	file, err := c.FormFile("file")
	if err==nil{
		if err = c.SaveUploadedFile(file, "images/"+file.Filename); err != nil {
			c.String(400, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}			
	}
}