package controllers

import (
	"fmt"
	"net/http"
	"tecnotree-go-project/entities"
	"tecnotree-go-project/interfaces"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService interfaces.IProduct
}

func InitProductController(productSvc interfaces.IProduct) *ProductController {
	return &ProductController{ProductService: productSvc}
}

func (p ProductController) AddProduct(c *gin.Context) {
	fmt.Println("Product Controller Invoked")

	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		return
	}
	result, err := p.ProductService.AddProduct(&product)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

func (p ProductController) SearchProducts(c *gin.Context) {
	fmt.Println("Products Controller Invoked")

	result, err := p.ProductService.SearchProducts()
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func (p ProductController) SearchProductById(c *gin.Context) {
	fmt.Println("Products Controller Invoked")
	var productId entities.SearchById
	err := c.BindJSON(&productId)
	if err != nil {
		return
	}
	result, err := p.ProductService.SearchProductById(&productId)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}
