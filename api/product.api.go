package api

import (
	"main/interceptor"
	"github.com/gin-gonic/gin"
)


// SetupProductAPI - call this method to setup product route group
func SetupTransactionAPI(router *gin.Engine) {
	productAPI:= router.Group("/api/v2")
	{
		productAPI.GET("/product" ,/*interceptor.GeneralInterceptor1*/interceptor.JwtVerify, getProduct)
		productAPI.POST("/product" /*interceptor.JwtVerify,*/, createProduct)
	}
}
func getProduct(c *gin.Context)  {
	c.JSON(200,gin.H{"result":"getProduct","username":c.GetString("jwt_username"),"Level":c.GetString("jwt_level")})
	//c.JSON(200,gin.H{"result":"getProduct"})
}
func createProduct(c *gin.Context)  {
	c.JSON(200,gin.H{"result":"createProduct"})
}