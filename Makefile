build:
	@go build -o bin/expensesapi cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/expensesapi
