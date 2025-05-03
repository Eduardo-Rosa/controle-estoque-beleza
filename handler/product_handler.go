package handler

import (
	"net/http"

	"controle-estoque-beleza/domain"

	"github.com/gin-gonic/gin"
)

// @Summary Lista todos os produtos
// @Description Retorna uma lista de produtos cadastrados
// @Tags Produtos
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Product
// @Router /api/v1/products [get]
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Lista de produtos"})
}

// @Summary Cria um novo produto
// @Description Adiciona um novo produto ao estoque
// @Tags Produtos
// @Accept  json
// @Produce  json
// @Param product body domain.Product true "Dados do produto"
// @Success 201 {object} domain.Product
// @Router /api/v1/products [post]
func CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Produto criado", "product": product})
}

// @Summary Atualiza um produto
// @Description Atualiza os dados de um produto existente
// @Tags Produtos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do produto"
// @Param product body domain.Product true "Dados do produto"
// @Success 200 {object} domain.Product
// @Router /api/v1/products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produto atualizado", "id": id, "product": product})
}

// @Summary Deleta um produto
// @Description Remove um produto do estoque
// @Tags Produtos
// @Param id path int true "ID do produto"
// @Success 204
// @Router /api/v1/products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusNoContent, gin.H{"message": "Produto deletado", "id": id})
}

// @Summary Busca produtos por categoria
// @Description Retorna uma lista de produtos filtrados por categoria
// @Tags Produtos
// @Param category query string true "Categoria do produto"
// @Success 200 {array} domain.Product
// @Router /api/v1/products/search [get]
func SearchProductsByCategory(c *gin.Context) {
	category := c.Query("category")
	c.JSON(http.StatusOK, gin.H{"message": "Produtos encontrados", "category": category})
}
