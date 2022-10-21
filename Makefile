install-hooks:
	bash ./hooks/install-hooks.sh 

live-reload:
	bash ./development-startup.sh
	air

start-dev:
	bash ./development-startup.sh 
	go run cmd/go-echo/server.go

start-stg:
	go run cmd/go-echo/server.go --env staging

start-prd:
	go run cmd/go-echo/server.go --env production

test:
	echo `go test -v ./...`

lint:
	golangci-lint run

build:
	go build -v -o server cmd/go-echo/server.go

air:
	go build -o ./tmp/main ./cmd/go-echo/server.go