mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/UnTea/WindingPacketsOnAPear/db/sqlc Store

initps:
	docker run --name postgres_db --network bank_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine3.18

startps:
	docker start postgres_db

stopps:
	docker stop postgres_db

createdb:
	docker exec -it postgres_db createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres_db dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src sqlc/sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

.PHONY: mockgen postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server
