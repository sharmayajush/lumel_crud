package service

import (
	"errors"
	"log"

	"github.com/sharmayajush/lumel_crud/utils/database"
)

type RevenueService struct{}

func (rev *RevenueService) GetTotalRevenue(start, end string) (interface{}, interface{}) {
	var totalRevenue float64
	if err := database.GetInstance().Table("orders").
		Joins("JOIN products ON orders.product_id = products.id").
		Where("orders.date_of_sale BETWEEN ? AND ?", start, end).
		Select("SUM(orders.quantity * (products.price - orders.discount) )").
		Scan(&totalRevenue).Error; err != nil {
		log.Println("unable to get total revenue from db: ", err)
		return nil, errors.New("unable to get total revenue from db")
	}
	return map[string]float64{"total_revenue": totalRevenue}, nil
}

func (rev *RevenueService) GetRevenueByProduct(productID, start, end string) (interface{}, interface{}) {

	var revenue float64
	err := database.GetInstance().Table("orders").
		Joins("JOIN products ON orders.product_id = products.id").
		Where("orders.product_id = ? AND orders.date_of_sale BETWEEN ? AND ?", productID, start, end).
		Select("SUM(orders.quantity * (products.price - orders.discount))").
		Scan(&revenue).Error
	if err != nil {
		log.Println("unable to get total revenue from db: ", err)
		return nil, errors.New("unable to get total revenue from db")
	}

	return map[string]float64{"revenue": revenue}, nil
}

func (rev *RevenueService) GetRevenueByCategory(category, start, end string) (interface{}, interface{}) {

	var revenue float64
	err := database.GetInstance().Table("orders").
		Joins("JOIN products ON orders.product_id = products.id").
		Where("products.category = ? AND orders.date_of_sale BETWEEN ? AND ?", category, start, end).
		Select("SUM(orders.quantity * (products.price - orders.discount))").
		Scan(&revenue).Error
	if err != nil {
		log.Println("unable to get total revenue from db: ", err)
		return nil, errors.New("unable to get total revenue from db")
	}

	return map[string]float64{"revenue": revenue}, nil
}

func (rev *RevenueService) GetRevenueByRegion(region, start, end string) (interface{}, interface{}) {

	var revenue float64
	err := database.GetInstance().Table("orders").
		Joins("JOIN products ON orders.product_id = products.id").
		Where("orders.region = ? AND orders.date_of_sale BETWEEN ? AND ?", region, start, end).
		Select("SUM(orders.quantity * (products.price - orders.discount))").
		Scan(&revenue).Error
	if err != nil {
		log.Println("unable to get total revenue from db: ", err)
		return nil, errors.New("unable to get total revenue from db")
	}

	return map[string]float64{"revenue": revenue}, nil
}
