# GRPC Microservices

This repo is based on https://github.com/huseyinbabal/microservices/tree/main

## Libraries Used

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## Calling GRPC 

```
grpcurl -d '{"user_id": 123, "items": [{"productCode": "abc123", "unit_price":100.0, "quantity":2}]}' -plaintext localhost:3000 Order/Create
```

