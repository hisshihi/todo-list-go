run:
	go run main.go

build:
	go build -o todo-list-go main.go

run-build:
	./todo-list-go

clean:
	rm -f todo-list-go

.PHONY:
	run build run-build clean