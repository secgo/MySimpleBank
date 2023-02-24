postgres:
	docker run --name banks -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
createdb:
	docker exec -it banks createdb --username=root --owner=root bank
dropdb:
	docker exec -it banks dropdb  bank

initMigrate:
	migrate create -ext sql -dir db/migrations -seq init_schema
upmigrate:
	migrate -path db/migrations/ -database "postgresql://root:secret@127.0.0.1/bank?sslmode=disable" up
downmigrate:
	migrate -path db/migrations/ -database "postgresql://root:secret@127.0.0.1/bank?sslmode=disable" down

initsqlc:
	sqlc init
generate:
	sqlc generate