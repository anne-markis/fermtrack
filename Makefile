.PHONY: start stop logs db

GOOSE_DRIVER=mysql
GOOSE_DBSTRING=root:s3CrEt@tcp(mysql:3306)/fermtrack?parseTime=true
GOOSE_MIGRATION_DIR=migrations

start:
	docker-compose up -d

stop:
	docker-compose down

test:
	go test ./... -v

logs:
	docker-compose logs -f fermtrack

db:
	docker exec -it fermtrack-mysql-1 mysql -uroot -ps3CrEt