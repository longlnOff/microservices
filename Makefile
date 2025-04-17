test_grpc:
	grpcurl -d '{"user_id": 123, "order_items": [{"product_code": "prod","quantity": 4, "unit_price": 12}]}' -plaintext localhost:3000 Order/Create

run_docker_mysql:
	docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=verysecretpass -e MYSQL_DATABASE=order mysql
