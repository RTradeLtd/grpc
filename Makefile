all: check

# Run simple checks
.PHONY: check
check:
	go vet ./...
	go test -run xxxx ./...

# Install dependencies
.PHONY: deps
deps:
	GO111MODULE=on go mod vendor
	bash .scripts/install-go-gen.sh

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

	make proto-service PKG=pay
	make proto-pkg PKG=pay/request
	make proto-pkg PKG=pay/response

	make proto-service PKG=nexus

	make proto-service PKG=krab

	make proto-service PKG=lensv2

.PHONY: proto-service
proto-service:
	protoc -I $(PKG) $(PKG)/service.proto --go_out=plugins=grpc:$(PKG)

.PHONY: proto-pkg
proto-pkg:
	protoc -I $(PKG) $(PKG)/*.proto --go_out=plugins=grpc:$(PKG)
