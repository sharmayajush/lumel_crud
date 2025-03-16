package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	contollers "github.com/sharmayajush/lumel_crud/src/controllers"
	"github.com/sharmayajush/lumel_crud/src/service"
	"github.com/sharmayajush/lumel_crud/utils/config"
	"github.com/sharmayajush/lumel_crud/utils/database"
)

var databaseService service.DatabaseService = service.DatabaseService{}

func init() {
	config.InitViper()
	database.GetInstance()
	database.InitDBModels()
	databaseService.ImportCSVToDB()
}

func main() {

	router := gin.Default()
	setupRoutes(router)
	go startServer(router)

	var wg sync.WaitGroup
	wg.Add(1)
	go refreshDatabase(&wg)
	wg.Wait()
}

func refreshDatabase(wg *sync.WaitGroup) {
	defer wg.Done()
	scheduler := gocron.NewScheduler(time.UTC).SingletonMode()
	_, err := scheduler.Every(1).Day().At("00:00").Do(databaseService.ImportCSVToDB)
	if err != nil {
		log.Printf("failed to configure cron task to reset database. err: %v", err)
		return
	}
	scheduler.StartBlocking()
}

func setupRoutes(router *gin.Engine) {
	fmt.Println("Setting up routes...")
	{
		db := router.Group("/db")
		{
			db.POST("/refresh", contollers.DBRefresh)
		}

		// Revenue routes
		revenue := router.Group("/revenue")
		{
			revenue.GET("/total", contollers.GetTotalRevenue)
			revenue.GET("/product/:product_id", contollers.GetRevenueByProduct)
			revenue.GET("/category/:category", contollers.GetRevenueByCategory)
			revenue.GET("/region/:region", contollers.GetRevenueByRegion)
		}
	}
}

func startServer(r *gin.Engine) {
	log.Println("Server started listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
