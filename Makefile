run:
	go run main.go

build:
	go build -o todo-list-go main.go

run-build:
	./todo-list-go

clean:
	rm -f todo-list-go

postgres:
	docker run --name todo-list-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=todo-list -p 5432:5432 -d postgres

createdb:
	docker exec -it todo-list-postgres createdb --username=postgres --owner=postgres todo-list

dropdb:
	docker exec -it todo-list-postgres dropdb --username=postgres todo-list

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo-list?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo-list?sslmode=disable" -verbose down

.PHONY:
	run build run-build clean createdb dropdb postgres migrateup migratedown