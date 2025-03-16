package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sharmayajush/lumel_crud/src/model"
	"github.com/sharmayajush/lumel_crud/utils/config"
	"github.com/sharmayajush/lumel_crud/utils/database"
	"github.com/spf13/viper"
)

func init() {
	config.InitViper()
	database.GetInstance()
	database.InitDBModels()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go refreshDatabase(&wg)
	wg.Wait()
}

func refreshDatabase(wg *sync.WaitGroup) {
	defer wg.Done()
	scheduler := gocron.NewScheduler(time.UTC).SingletonMode()
	_, err := scheduler.Every(1).Day().At("00:00").Do(importCSVToDB)
	if err != nil {
		log.Printf("failed to configure cron task to reset database. err: %v", err)
		return
	}
	scheduler.StartBlocking()
}

func importCSVToDB() {
	log.Println("Running cron job to import CSV to db at: ", time.Now())

	db := database.GetInstance()

	file, err := os.Open(viper.GetString("csv.path"))
	if err != nil {
		log.Println("error opening file: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("error reading csv: ", err)
		return
	}
	for _, row := range records[1:] {
		orderID, prodID, custID, prodName, category, region, saleDate, quantity, price, discount, shippingCost, PaymentMethod, custName, custEmail, custAddr := row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14]
		var id string
		if err := db.Model(&model.Customer{}).Select("id").Where("id = ?", custID).Scan(id).Error; err != nil {
			log.Printf("unable to fetch customer id: %s", err.Error())
			return
		}
		if id != "" {
			if err := db.Where("id = ?", custID).UpdateColumns(&model.Customer{
				Name:    custName,
				Email:   custEmail,
				Address: custAddr,
			}).Error; err != nil {
				fmt.Println("unable to update customer table:", err)
				return
			}
		} else {
			if err := db.Create(&model.Customer{
				ID:      custID,
				Name:    custName,
				Email:   custEmail,
				Address: custAddr,
			}).Error; err != nil {
				fmt.Println("unable to create customer table:", err)
				return
			}
		}

		floatPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("error converting price to float:", err)
			return
		}
		id = ""
		if err := db.Model(&model.Product{}).Select("id").Where("id = ?", prodID).Scan(id).Error; err != nil {
			log.Printf("unable to fetch product id: %s", err.Error())
			return
		}
		if id != "" {
			if err := db.Where("id = ?", prodID).UpdateColumns(&model.Product{
				Name:     prodName,
				Category: category,
				Price:    floatPrice,
			}).Error; err != nil {
				fmt.Println("unable to update product table:", err)
				return
			}

		} else {
			if err := db.Create(&model.Product{
				ID:       prodID,
				Name:     prodName,
				Category: category,
				Price:    floatPrice,
			}).Error; err != nil {
				fmt.Println("unable to create product table:", err)
				return
			}
		}
		intOrderID, _ := strconv.ParseUint(orderID, 10, 64)
		intQuantity, _ := strconv.ParseInt(quantity, 10, 64)
		floatDiscount, _ := strconv.ParseFloat(discount, 64)
		floatShip, _ := strconv.ParseFloat(shippingCost, 64)
		date, _ := time.Parse("2006-01-02", saleDate)
		id = ""
		if err := db.Model(&model.Order{}).Select("id").Where("id = ?", orderID).Scan(id).Error; err != nil {
			log.Printf("unable to fetch product id: %s", err.Error())
			return
		}
		if id != "" {
			if err := db.Where("id = ?", orderID).UpdateColumns(&model.Order{
				CustomerID:    custID,
				ProductID:     prodID,
				Quantity:      int(intQuantity),
				Discount:      floatDiscount,
				ShippingCost:  floatShip,
				PaymentMethod: PaymentMethod,
				DateOfSale:    date,
				Region:        region,
			}).Error; err != nil {
				fmt.Println("unable to update orders table:", err)
				return
			}

		} else {
			if err := db.Create(&model.Order{
				ID:            intOrderID,
				CustomerID:    custID,
				ProductID:     prodID,
				Quantity:      int(intQuantity),
				Discount:      floatDiscount,
				ShippingCost:  floatShip,
				PaymentMethod: PaymentMethod,
				DateOfSale:    date,
				Region:        region,
			}).Error; err != nil {
				fmt.Println("unable to insert data in orders table:", err)
				return
			}

		}

	}

}
