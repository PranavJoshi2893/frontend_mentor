include .env
export

build:
	@go build -o ./bin/frontend_mentor_v1 ${PWD}/cmd/main.go

run: build
	@./bin/frontend_mentor_v1

clean:
	@go clean
	@rm -r ./bin