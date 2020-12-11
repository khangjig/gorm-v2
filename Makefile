.PHONY: build local-db proto run unit-test open-coverage lint lint-consisten

build:
	go build -ldflags "-s -w" -o ./tmp/server ./cmd/main.go

local-db:
	@docker-compose down
	@docker-compose up -d

proto:
	@protoc --go_out=plugins=grpc:./proto --proto_path=proto --go_opt=paths=source_relative proto/service.proto
	sed -i "" -e "s/,omitempty//g" ./proto/*.go # remove 'omitempty'

unit-test:
	@mkdir coverage || true
	@go test -race -v -coverprofile=coverage/coverage.txt.tmp -count=1  ./...
	@cat coverage/coverage.txt.tmp | grep -v "mock_" > coverage/coverage.txt
	@go tool cover -func=coverage/coverage.txt
	@go tool cover -html=coverage/coverage.txt -o coverage/index.html

open-coverage:
	@open coverage/index.html

run:
	@GO111MODULE=off go get -u github.com/cosmtrek/air
	@air -c .air.conf

lint:
	@hash golangci-lint 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0
	@GO111MODULE=on CGO_ENABLED=0 golangci-lint run

lint-consistent:
	@hash go-consistent 2>/dev/null || GO111MODULE=off go get -v github.com/quasilyte/go-consistent
	@go-consistent ./...
