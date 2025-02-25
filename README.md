# Создание миграций

```bash
migrate create -ext sql -dir db/migration -seq init_schema

create - создает новую миграцию
ext - расширение файла
dir - директория с миграциями
seq - название миграции
```

## Применение миграций

```bash
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/todo-list?sslmode=disable" -verbose up
```
