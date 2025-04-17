test_grpc:
	grpcurl -d '{"user_id": 123, "order_items": [{"product_code": "prod","quantity": 4, "unit_price": 12}]}' -plaintext localhost:3000 Order/Create

run_docker_mysql:
	docker run --name mysql-microservice -p 3306:3306 -e MYSQL_ROOT_PASSWORD=verysecretpass -e MYSQL_DATABASE=order -d mysql 
create_db:
	docker exec mysql-microservice mysql -uroot -pverysecretpass -e "CREATE DATABASE payment; CREATE DATABASE shipping;"
start_docker_mysql:
	docker start mysql-microservice
stop_docker_mysql:
	docker stop mysql-microservice
delete_docker_mysql:
	docker rm mysql-microservice

service_payment:
	cd payment && DB_DRIVER=mysql DATA_SOURCE_URL=root:verysecretpass@tcp\(127.0.0.1:3306\)/payment APPLICATION_PORT=3001 ENV=development go run cmd/main.go

service_order:
	cd order && DB_DRIVER=mysql DATA_SOURCE_URL=root:verysecretpass@tcp\(127.0.0.1:3306\)/order APPLICATION_PORT=3000 ENV=development PAYMENT_SERVICE_URL=localhost:3001 go run cmd/main.go
