build:
	go build -o bin/app main.go

run:
	go run main.go

client:
	go run ./client/main.go

watch:
	templ generate --watch --proxy="http://localhost:3000" --cmd="go run ."

.PHONY: watch