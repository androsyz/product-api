# product-api
[![Generic badge](https://img.shields.io/badge/go-v1.20-blue.svg)](https://shields.io/) 

Restful API that provides endpoints to manage products. Build with Gin Gonic and Gorm.

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [RESTful API Endpoints](#restful-api-endpoints)
* [Setup](#setup)

## General info
This app built with Gin Gonic and Gorm, the database used is MySQL

### Features
* CRUD:
  - [x]  The API enables clients to perform the full set of CRUD operations on products, including creating, reading, updating, and deleting records.
 
* Filtering:
  - [x]  The GET /api/products endpoint allows clients to filter the list of products by category using the category query parameter. 
  
* Pagination:
  - [x]  The GET /api/products endpoint returns a paginated list of products by default, with five products per page. Clients can control the page number and page size using the page and page_size query parameters, respectively.


## Technologies
Project is created with:
* go : 1.20
* gin-gonic : 1.9.0
* gorm : 1.24.5

## RESTful API Endpoints
### API Endpoints
RESTful API Endpoints are shown in the table below:
| Method | Endpoint | Description |
| --- | --- | --- |
| POST | `/api/products` |  Create a new product | 
| GET | `/api/products` | Retrieve a list of products |
| GET | `/api/products/{id}`|  Retrieve a single product by ID |
| PUT | `/api/products/{id}` | Update a product by ID|
| DELETE | `/api/products/{id}` | Delete a product by ID |
### Parameters
* POST & PUT :
#### Parameter :
| Name | Type | Description |
| --- | --- | --- |
| product_name | string | The name of the product |
| category | string | The category of the product |
| description | string | The description of the product. |
#### Example Request :
```json
{ 
  "product_name" : "product", 
  "category" : "category", 
  "description" : "description"
}
```
* GET :
#### Query Parameter :
| Name | Type | Description |
| --- | --- | --- |
| category | string | The category of the products to retrieve |
| page | integer | The page number to retrieve |
| page_size | integer | The number of products to retrieve per page (default 5) |
#### Example Endpoint :
```bash
GET /api/products?category=fashion&page=1&page_size=5
```



## Setup
To run this project, please install it locally:
```
$ cd product-api
$ go get
```
please add .env in root folder before run the application, the suggested .env is bellow:
```
DB_HOST="127.0.0.1"
DB_PORT="3306"
DB_USER="your_user_name"
DB_NAME="your_db_name"
DB_PASSWORD="your_db_password"
```
You can run the app by entering code bellow:
```
$ go run main.go
```
