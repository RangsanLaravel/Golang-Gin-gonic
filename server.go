package main

import ("github.com/gin-gonic/gin"
		"main/api")

func main()  {
	r := gin.Default()
	r.Static("/images","./uploaded/images")
	api.Setup(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
