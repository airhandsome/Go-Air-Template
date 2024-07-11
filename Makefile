.PHONY: build run test docker generate

APP := myapp
IMAGE := my-go-project

build:
	@go build -o ./bin/$(APP) ./cmd/$(APP)

run:
	@./bin/$(APP)

test:
	@go test ./...

docker:
	@docker build -t $(IMAGE) .
	@docker run -it --rm $(IMAGE)

generate:
	@./scripts/generate.sh

clean:
	@rm -rf ./bin
	@docker rmi $(IMAGE) -f
lint:   
	@golangci-lint run --timeout 5m
sql-xo: 
	@xo ./sql/psql/*.sql -o internal/domain
