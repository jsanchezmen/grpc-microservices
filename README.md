# GRPC Microservices

This repo is based on https://github.com/huseyinbabal/microservices/tree/main

## Libraries Used

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
# Order Folder
go get -u github.com/jsanchezmen/microservices-proto/golang/order@v0.0.1-beta
# Payment Folder
go get -u github.com/jsanchezmen/microservices-proto/golang/payment@v0.0.2-beta
```

## Calling GRPC 

```
grpcurl -d '{"user_id": 123, "items": [{"productCode": "abc123", "unit_price":100.0, "quantity":2}]}' -plaintext localhost:3000 Order/Create
grpcurl -d '{"user_id": 123, "order_id": 111, "total_price": 222}' -plaintext 127.0.0.1:3001 Payment/Create
```

## Runnig MySql container

```
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=microservices mysql
# Exec container
docker exec -it c37163af0621 bash
# Enter MySql
mysql -u root -p
show databases;
use microservices;
show tables;
```

## Environment Variables
```
#export DATA_SOURCE_URL={user}:{password}@tcp(127.0.0.1:3306)/{database}
export DATA_SOURCE_URL=root:root@tcp(127.0.0.1:3306)/microservices
export ENV=dev
export ORDER_APPLICATION_PORT=3000
export PAYMENT_APPLICATION_PORT=3001
```

## Run app

```
go run cmd/main.go
```