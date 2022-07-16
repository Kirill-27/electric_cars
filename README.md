# electric_cars
migrate create -ext sql -dir ./schema -seq init \n
psql postgresql://cars:cars@localhost:5435 \n
migrate -path ./schema -database 'postgresql://cars:cars@localhost:5435/cars?sslmode=disable' up \n