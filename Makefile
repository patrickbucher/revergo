.PHONY: test coverage vet lint

test:
	go test ./...

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm -f coverage.out

vet:
	go vet ./...

lint:
	golint ./...
