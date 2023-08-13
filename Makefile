postgres:
	docker run --name postgres_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine3.18

createdb:
	docker exec -it postgres_db createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres_db dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src sqlc/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc