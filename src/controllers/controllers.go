package contollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharmayajush/lumel_crud/src/service"
)

var databaseService service.DatabaseService = service.DatabaseService{}
var revenueService service.RevenueService = service.RevenueService{}

func DBRefresh(c *gin.Context) {
	insert, err := databaseService.ImportCSVToDB()
	if err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}
	respondWithSuccess(c, http.StatusOK, insert)
}

func GetTotalRevenue(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	res, err := revenueService.GetTotalRevenue(startDate, endDate)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithSuccess(c, http.StatusOK, res)
}

func GetRevenueByProduct(c *gin.Context) {
	productID := c.Param("product_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	res, err := revenueService.GetRevenueByProduct(productID, startDate, endDate)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithSuccess(c, http.StatusOK, res)
}

func GetRevenueByCategory(c *gin.Context) {
	category := c.Param("category")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	res, err := revenueService.GetRevenueByCategory(category, startDate, endDate)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithSuccess(c, http.StatusOK, res)
}

func GetRevenueByRegion(c *gin.Context) {
	region := c.Param("region")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	res, err := revenueService.GetRevenueByRegion(region, startDate, endDate)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithSuccess(c, http.StatusOK, res)
}

// Response Handling Functions
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{"error": message})
}

func respondWithSuccess(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{"response": message})
}
