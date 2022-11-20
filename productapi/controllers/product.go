package controllers

import (
	"ilmudata/productapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type ProductRepo struct {
	Products *[] models.Product
}
func Init(products *[]models.Product) *ProductRepo {
	return &ProductRepo{Products: products}
}
// CRUD operations for controllers
func (repo *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Product

	c.BindJSON(&product)
	err := models.CreateProduct(repo.Products, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (repo *ProductRepo) ReadProducts(c *gin.Context) {	
	c.JSON(http.StatusOK, repo.Products)
}

func (repo *ProductRepo) ReadProductById(c *gin.Context) {	
	id, _ := strconv.ParseInt(c.Param("id"),10,0)

	product := models.ReadProductById(repo.Products, int(id))
	c.JSON(http.StatusOK, product)
}

func (repo *ProductRepo) UpdateProduct(c *gin.Context) {	
	id, _ := strconv.ParseInt(c.Param("id"),10,0)
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	var product models.Product
	c.BindJSON(&product)	
	if product.Id != int(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	udaptedProduct := models.UpdateProductById(repo.Products, &product)
	c.JSON(http.StatusOK, &udaptedProduct)
}

func (repo *ProductRepo) DeleteProductById(c *gin.Context) {	
	id, _ := strconv.ParseInt(c.Param("id"),10,0)

	product := models.DeleteProductById(repo.Products, int(id))
	c.JSON(http.StatusOK, product)
}
