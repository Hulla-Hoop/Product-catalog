migrate-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres dbname=test password=12345678 port=5432 sslmode=disable" goose up 
run:
	@docker-compose up -d

