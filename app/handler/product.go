package handler

import (
	"net/http"
	"onlineshop/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Create product
// @Security ApiKeyAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param name formData string true "product name"
// @Param image formData file true "poster"
// @Param price formData string true "price"
// @Param height formData string true "height"
// @Param size formData string true "size"
// @Param instruction formData string false "instruction"
// @Param description formData string true "description"
// @Param recommended_products formData []string false "recommended products"
// @Success 200 {integer} integer
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/admin/products [post]
func (h *Handler) createProduct(c *gin.Context) {

	var input models.Product

	name := c.PostForm("name")
	image, err := c.FormFile("image")
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	price := c.PostForm("price")
	pricee, err := strconv.ParseFloat(price, 64)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	path := "files//product_images//" + image.Filename
	c.SaveUploadedFile(image, path)
	height := c.PostForm("height")
	size := c.PostForm("size")
	instruction := c.PostForm("instruction")
	description := c.PostForm("description")
	recommendedProducts := c.PostFormArray("recommended_products")

	logrus.Print(recommendedProducts)

	input = models.Product{
		Name:                name,
		Image:               path,
		Price:               pricee,
		Height:              height,
		Size:                size,
		Instruction:         instruction,
		Description:         description,
		RecommendedProducts: recommendedProducts}

	id, err := h.services.Product.Create(input)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary All products
// @Security ApiKeyAuth
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} models.GetProducts
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/products [get]
func (h *Handler) getAlProdacts(c *gin.Context) {

	products, err := h.services.Product.GetAll()
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Data": products,
	})

}

// @Summary delete product
// @Security ApiKeyAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/admin/products/{id} [delete]
func (h *Handler) deleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Product.Delete(productId); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary edit the product
// @Security ApiKeyAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Param name formData string false "product name"
// @Param image formData file false "poster"
// @Param price formData string false "price"
// @Param height formData string false "height"
// @Param size formData string false "size"
// @Param instruction formData string false "instruction"
// @Param description formData string false "description"
// @Param recommended_products formData []string false "recommended products"
// @Success 200 {integer} integer
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/admin/products/{id} [put]
func (h *Handler) updateProduct(c *gin.Context) {

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	name := c.PostForm("name")

	image, err := c.FormFile("image")
	var path string
	if err != nil {
		product, _, err := h.services.Product.GetById(productId)
		if err != nil {
			NewErrorResponce(c, http.StatusInternalServerError, err.Error()+"getProduct")
			return
		}
		path = product.Image
	} else {
		c.SaveUploadedFile(image, path)
		path = "files//product_images//" + image.Filename
	}
	price := c.PostForm("price")
	var pricee float64
	if price != "" {
		priceee, err := strconv.ParseFloat(price, 64)
		if err != nil {
			NewErrorResponce(c, http.StatusBadRequest, err.Error()+"strcon")
			return
		}
		pricee = priceee

	}

	height := c.PostForm("height")
	size := c.PostForm("size")
	instruction := c.PostForm("instruction")
	description := c.PostForm("description")
	recommendedProducts := c.PostFormArray("recommended_products")

	input := models.UpdateProduct{
		Name:                name,
		Image:               path,
		Price:               pricee,
		Height:              height,
		Size:                size,
		Instruction:         instruction,
		Description:         description,
		RecommendedProducts: recommendedProducts}

	logrus.Print(input.Instruction)

	if err := h.services.Product.Update(productId, input); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary product by id
// @Security ApiKeyAuth
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/products/{id} [get]
func (h *Handler) getProdactById(c *gin.Context) {

	prodactId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	product, products, err := h.services.Product.GetById(prodactId)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Product":              product,
		"Recommended products": products,
	})

}
