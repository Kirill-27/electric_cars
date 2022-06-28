# electric_cars
migrate create -ext sql -dir ./schema -seq init
psql postgresql://cars:cars@localhost:5435
migrate -path ./schema -database 'postgresql://cars:cars@localhost:5435/cars?sslmode=disable' up