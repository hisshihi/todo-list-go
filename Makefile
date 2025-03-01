server:
	go run main.go

test:
	go test -v -cover -race -timeout 30s -count=1 ./...

postgres:
	docker run --name todo-list-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=todo-list -p 5432:5432 -d postgres

createdb:
	docker exec -it todo-list-postgres createdb --username=postgres --owner=postgres todo-list

dropdb:
	docker exec -it todo-list-postgres dropdb --username=postgres todo-list

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo-list?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo-list?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo-list?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo-list?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

.PHONY:
	server postgres createdb dropdb migrateup migratedown migratedown1 migrateup1 sqlc