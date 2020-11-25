package api

import "github.com/gin-gonic/gin"

func SetupTransactionAPI(router *gin.Engine) {
	productAPI:= router.Group("/api/v2")
	{
		productAPI.GET("/product" /*interceptor.JwtVerify,*/, getProduct)
		productAPI.POST("/product" /*interceptor.JwtVerify,*/, createProduct)
	}
}
func getProduct(c *gin.Context)  {
	c.JSON(200,gin.H{"result":"getProduct"})
}
func createProduct(c *gin.Context)  {
	c.JSON(200,gin.H{"result":"createProduct"})
}