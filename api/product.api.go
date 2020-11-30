package api

import (
	"fmt"
	"main/db"
	"main/interceptor"
	"main/model"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupProductAPI - call this method to setup product route group
func SetupTransactionAPI(router *gin.Engine) {
	productAPI:= router.Group("/api/v2")
	{
		productAPI.GET("/product" ,/*interceptor.GeneralInterceptor1*/interceptor.JwtVerify, getProduct)
		productAPI.POST("/product" ,/*interceptor.JwtVerify,*/ createProduct)
		productAPI.PUT("/product" ,/*interceptor.JwtVerify,*/ editProduct)
	}
}
func getProduct(c *gin.Context)  {
	c.JSON(200,gin.H{"result":"getProduct","username":c.GetString("jwt_username"),"Level":c.GetString("jwt_level")})
	//c.JSON(200,gin.H{"result":"getProduct"})
}
func editProduct(c *gin.Context) {
	var product model.Product
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)

	product.ID = uint(id)
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CreatedAt = time.Now()
	db.GetDB().Updates(&product)

	image, _ := c.FormFile("image")
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})
}

func createProduct(c *gin.Context)  {
	product:=model.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CreatedAt = time.Now()
	image,_:=c.FormFile("image")
	 db.GetDB().Create(&product)
	 saveImage(image, &product, c)
	 c.JSON(http.StatusOK, gin.H{"result": product})
	//c.JSON(200,gin.H{"result":"getProduct"})
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	if image != nil {
		runningDir, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, fileName)

		if fileExists(filePath) {
			os.Remove(filePath)
		}
		c.SaveUploadedFile(image, filePath)
		db.GetDB().Model(&product).Update("image", fileName)
	}
}