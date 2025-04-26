# Product Management API

## Overview
RESTful API for Product Management

## Tech Stack
- Go 1.20, Gin, GORM  
- PostgreSQL  
- Docker, docker-compose  

## Setup

## Setup with Docker Compose

### 1. Clone the repository
git clone github.com/ltvinh9899/soa_test.git
cd soa_test

### 2. Create envỉonment variables (.env)

PORT=8080
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=online_store
JWT_SECRET=your_jwt_secret

### 3. Build and start containers
docker-compose up --build -d

### 5.  Initial categories
docker-compose exec db psql -U postgres -d online_store -c "
INSERT INTO categories (name, description, created_at, updated_at) VALUES
('Electronics', 'Electronic devices', NOW(), NOW()),
('Accessories', 'Various accessories', NOW(), NOW()),
('Laptops', 'Laptop computers', NOW(), NOW()),
('Monitors', 'Computer displays', NOW(), NOW());

### API Documentation
Authentication
Method	Endpoint	Description	Access
POST	api/user/login	User login	Public
POST	api/user/register	User registration	Public

Products
Method	Endpoint	Description	Access
GET	/api/products	List products	Public
POST	/api/product	Create product	Admin
GET	/api/product/{id}	Get product	Public
PUT	/api/product/{id}	Update product	Admin
DELETE	/api/products/{id}	Delete product	Admin

Dashboard
Method	Endpoint	Description	Access
GET	/api/dashboard	Category stats	Admin