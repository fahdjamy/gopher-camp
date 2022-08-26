setup: local
	. ./local/setup.sh

run-dev:
	#go build cmd/main.go -o bin/app && ./bin/app
	go build cmd/main.go && go run cmd/main.go
