package customerHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kisahtegar/bootcamp-api-hmsi/models"
	"github.com/kisahtegar/bootcamp-api-hmsi/modules/customers"
)

type customerHandler struct {
	UC customers.CustomerUsecase
}

func NewCustomerHandler(r *gin.Engine, UC customers.CustomerUsecase) {
	handler := customerHandler{UC}

	r.GET("/customers", handler.GetAll)
	r.POST("/customer/create", handler.Insert)
	r.PUT("/customer/update", handler.Update)
	r.DELETE("/customer/delete/:id", handler.Delete)
}

// Implement handler GetAll customer.
func (h *customerHandler) GetAll(c *gin.Context) {
	results, err := h.UC.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "success",
		"data":    results,
	})
}

// Implement handler Insert customer.
func (h *customerHandler) Insert(c *gin.Context) {
	var request models.RequestInsertCustomer

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})

		return
	}

	err := h.UC.Insert(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	// When not error then...
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Inserted successfully",
		"data":    []string{},
	})
}

// Implement handler Update customer.
func (h *customerHandler) Update(c *gin.Context) {
	var request models.RequestUpdateCustomer

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	err := h.UC.Update(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	// When not error then...
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Updated successfully",
		"data":    []string{},
	})
}

// Implement handler Delete customer.
func (h *customerHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	// Check if id empty or not
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID cannot be empty",
			"data":    []string{},
		})
		return
	}

	// Convert string to uint64
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID format",
			"data":    []string{},
		})
		return
	}

	// Call Delete usecase.
	err = h.UC.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []string{},
		})
		return
	}

	// When not error then...
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted successfully",
		"data":    []string{},
	})
}
