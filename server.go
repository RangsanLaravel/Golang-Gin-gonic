package main

import ("github.com/gin-gonic/gin"
		"main/api")

func main()  {
	r := gin.Default()
	//set public folder
	r.Static("/images","./uploaded/images")
	//GO call api center
	api.Setup(r)
	//Test run
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
