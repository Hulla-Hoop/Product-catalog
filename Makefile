migrate-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5432 sslmode=disable" goose -dir ./migrations/ up 
migrate-down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5432 sslmode=disable" goose -dir ./migrations/ down 
run:
	@docker-compose up -d

