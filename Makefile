# Golang Makefile for package golang-stress-tester

VERSION ?= v0.1.0
PACKAGE_NAME ?= $(shell grep "^module " go.mod | sed 's/^module\s//')
RUN_ARGS ?= -h
BUILD_MAIN_FILE = main.go
GO_CMD = go
GOPATH ?= $(shell $(GO_CMD) env GOPATH)
GOAUTODOC_CMD = $(GOPATH)/bin/autodoc
ECHO = "/usr/bin/echo"
# GOAUTODOC_CMD = $(shell go env GOPATH)/bin/goautodoc
TIME = $(shell date)
# Linker-added variables (residing in the config/ module)
LDFLAGS = -X '$(PACKAGE_NAME)/config.PackageName=$(PACKAGE_NAME)' -X '$(PACKAGE_NAME)/config.Version=$(VERSION)' -X '$(PACKAGE_NAME)/config.BuildTime=$(TIME)'

test:
	go test utils/*.go
	go test session/*.go

get-pkg:
	go get gopkg.in/yaml.v2
	go get github.com/eschao/config
	go mod tidy

setup:
	go mod tidy
	$(GOAUTODOC_CMD) -version || go install github.com/projectbadger/autodoc@latest

build-linux-amd64:
	echo "Compiling linux-amd64"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/amd64/$(PACKAGE_NAME)-linux-amd64 $(BUILD_MAIN_FILE)

build-linux-arm64:
	echo "Compiling linux-arm64"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/arm64/$(PACKAGE_NAME)-linux-arm64 $(BUILD_MAIN_FILE)

build-linux-arm:
	echo "Compiling linux-arm"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/linux/arm/$(PACKAGE_NAME)-linux-arm $(BUILD_MAIN_FILE)

build-windows-amd64:
	echo "Compiling linux-arm64"
	CGO_ENABLED=0 GOOS=windows $(GO_CMD) build -ldflags="$(LDFLAGS)" -o release/windows/amd64/$(PACKAGE_NAME)-windows-amd64 $(BUILD_MAIN_FILE)

run:
	echo "Running the package..."
	$(GO_CMD) run -ldflags="$(LDFLAGS)" $(BUILD_MAIN_FILE) $(RUN_ARGS)

docs:
	$(GOAUTODOC_CMD) -config ./.autodoc/config.root.yml > ./README.md
	$(GOAUTODOC_CMD) -package ./config > ./config/README.md
	$(GOAUTODOC_CMD) -package ./log > ./log/README.md
	$(GOAUTODOC_CMD) -package ./utils > ./utils/README.md
	$(GOAUTODOC_CMD) -package ./rest > ./rest/README.md
	$(GOAUTODOC_CMD) -package ./rest/api_interface > ./rest/api_interface/README.md
	$(GOAUTODOC_CMD) -package ./rest/errors > ./rest/errors/README.md
	$(GOAUTODOC_CMD) -package ./rest/running > ./rest/running/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/access_gateway > ./rest/running/access_gateway/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/chassis > ./rest/running/chassis/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/fabric > ./rest/running/fabric/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/fdmi > ./rest/running/fdmi/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/fibrechannel_configuration > ./rest/running/fibrechannel_configuration/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/fibrechannel_diagnostics > ./rest/running/fibrechannel_diagnostics/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/brocade_interface > ./rest/running/brocade_interface/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/fibrechannel_logical_switch > ./rest/running/fibrechannel_logical_switch/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/fibrechannel_switch > ./rest/running/fibrechannel_switch/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/fibrechannel_trunk > ./rest/running/fibrechannel_trunk/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/fru > ./rest/running/fru/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/license > ./rest/running/license/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/logging > ./rest/running/logging/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/maps > ./rest/running/maps/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/media > ./rest/running/media/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/name_server > ./rest/running/name_server/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/operations > ./rest/running/operations/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/security > ./rest/running/security/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/snmp > ./rest/running/snmp/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/time > ./rest/running/time/README.md
	$(GOAUTODOC_CMD) -package ./rest/running/zone > ./rest/running/zone/README.md
	$(GOAUTODOC_CMD) -package ./session > ./session/README.md


info:
	$(ECHO) "Package:    '${PACKAGE_NAME}'" > /dev/null
	$(ECHO) "Version:    '${VERSION}'" > /dev/null
	$(ECHO) "Build file: '${BUILD_MAIN_FILE}'" > /dev/null
	$(ECHO) "Go command: '${GO_CMD}'" > /dev/null
	$(ECHO) "Time:       '$(TIME)'" > /dev/null

build-all: build-linux-amd64 build-linux-arm64 build-linux-arm build-windows-amd64
