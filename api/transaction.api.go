package api

import "github.com/gin-gonic/gin"

func SetupAuthenAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", createTransaction)
	}
}
func getTransaction(c *gin.Context)  {
	c.JSON(200,gin.H{"result":"getTransaction"})
}
func createTransaction(c *gin.Context)  {
	c.JSON(200,gin.H{"result":"createTransaction"})
}