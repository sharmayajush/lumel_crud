# lumel_crud

## Overview

The Revenue API provides endpoints to fetch revenue-related data such as total revenue, revenue by product, category, and region. The API is built using the Gin framework in Go and follows RESTful principles.

## Design Decisions

### 1. Gin Framework

- Used Gin for better performance and simplicity.

- Provides built-in middleware support for logging, recovery, and routing.

### 2. Database Design

- Uses a relational database with tables orders, customers and products.

- orders table contains sales data including product_id, quantity, date_of_sale, and discount.

- products table contains product details including id, name, and price.

- customers table containe customer specific details

### 3. Cron Job for Data Refresh

- Uses gocron to periodically refresh data by importing a CSV file.

- The scheduled job runs at midnight daily (00:00 UTC).

### 4. Error Handling

- API responses follow a structured format with success and error messages.

- HTTP status codes are used to indicate success (200 OK) and failures (400, 500).

### 5. Routing Structure

- All routes are prefixed for organization:

- /db/refresh for database-related operations.

- /revenue/... for revenue calculations.

- Path parameters (:product_id, :category, :region) for fetching specific revenue data.

## Schema Design

![alt text](https://github.com/sharmayajush/lumel_crud/blob/main/lumel_ss.png?raw=true)

## Code

### Prerequisites
Golang v1.23.6 and above, Postgresql with schema 'lumel', csv file
### Code Execution
- clone the repo github.com/sharmayajush/lumel_crud
```sh
git clone https://github.com/sharmayajush/lumel_crud.git
cd lumel_crud
```
- run command 'go run main.go'
```sh
go mod tidy
go run main.go
```
- this will start the go program and will start the api server and cron job for database refreshing.

## API Documentation

### List of APIs

| Route                        | Method | Request Body (if applicable) | Sample Response | Description |
|------------------------------|--------|------------------------------|-----------------|-------------|
| `/db/refresh`                | POST   | N/A                          | `{ "response": "Database refreshed successfully" }` | Refreshes the database by importing CSV data. |
| `/revenue/total`             | GET    | N/A (Query params required)  | `{ "response": { "total_revenue": 12345.67 } }` | Fetches total revenue between `start_date` and `end_date`. |
| `/revenue/product/:product_id` | GET  | N/A (Query params required)  | `{ "response": { "revenue": 5678.90 } }` | Fetches revenue for a specific product in a date range. |
| `/revenue/category/:category` | GET  | N/A (Query params required)  | `{ "response": { "revenue": 7890.12 } }` | Fetches revenue for a specific category in a date range. |
| `/revenue/region/:region`     | GET  | N/A (Query params required)  | `{ "response": { "revenue": 3456.78 } }` | Fetches revenue for a specific region in a date range. |

#### Query Parameters
Most of the revenue APIs require the following query parameters:
- **`start_date`** (YYYY-MM-DD) - Start date for filtering revenue.
- **`end_date`** (YYYY-MM-DD) - End date for filtering revenue.

### Example Usage
#### Get Total Revenue
```sh
curl -v "http://127.0.0.1:8080/revenue/total?start_date=2024-03-01&end_date=2024-03-15"
```
#### Get Revenue by Product
```sh
curl -v "http://127.0.0.1:8080/revenue/product/P456?start_date=2024-03-01&end_date=2024-03-15"
```

## Future Improvements

- Implement authentication & authorization.

- Add pagination support for large revenue reports.

- Optimize database queries for better performance.
