install-hooks:
	bash ./hooks/install-hooks.sh 

live-reload:
	bash ./development-startup.sh
	air

start-dev:
	bash ./development-startup.sh 
	go run server.go

start-stg:
	go run server.go --env staging

start-prd:
	go run server.go --env production

test:
	echo `go test -v ./...`

lint:
	golangci-lint run

build:
	go build -v -o server server.go
