all: check

# Run simple checks
.PHONY: check
check:
	go vet ./...
	go test -run xxxx ./...

# Install dependencies
.PHONY: deps
deps:
	dep ensure

# Execute tests
.PHONY: test
test:
	go test -race -cover ./...

# Run linters and checks
.PHONY: lint
lint: check
	go fmt ./...
	golint `go list ./... | grep -v /vendor/`

# Generate protobuf code from definitions
.PHONY: proto
proto:
	make proto-service PKG=lens
	make proto-pkg PKG=lens/request
	make proto-pkg PKG=lens/response

	make proto-service PKG=temporal
	make proto-pkg PKG=temporal/request
	make proto-pkg PKG=temporal/response

.PHONY: proto-service
proto-service:
	( cd $(PKG) ; protoc *.proto --go_out=. )

.PHONY: proto-pkg
proto-pkg:
	( cd $(PKG) ; protoc *.proto --go_out=plugins=grpc:$(GOPATH)/src )
