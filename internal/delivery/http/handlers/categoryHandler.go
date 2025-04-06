package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(u *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{*u}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req domain.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	category := &domain.Category{
		Name: req.Name,
	}
	if err := h.CategoryUseCase.CreateCategory(c.Request.Context(), category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"category": category})
}

func (h *CategoryHandler) GetCategoryById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Category ID"})
	}
	category, err := h.CategoryUseCase.GetCategoryByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "category not found"})
	}
	c.JSON(http.StatusOK, gin.H{"category": category})

}
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	category.ID = uint(id)

	updatedCategory, err := h.CategoryUseCase.UpdateCategory(c.Request.Context(), &category)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
	filter := make(map[string]interface{})

	if categoryID := c.Query("category_id"); categoryID != "" {
		if id, err := strconv.ParseUint(categoryID, 10, 64); err == nil {
			filter["category_id"] = uint(id)
		}
	}

	categories, err := h.CategoryUseCase.ListCategories(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	err = h.CategoryUseCase.DeleteCategory(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
