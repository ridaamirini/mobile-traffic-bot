OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

CMD_FOLDER = cmd
MAIN_FILE = main.go
ARGS = $(filter-out all subtest1, $(MAKECMDGOALS))

.PHONY: build

run-%:
	@go run $(CMD_FOLDER)/$(firstword $(subst _, ,$*))/$(MAIN_FILE) --account="$(account)"

build-%:
	@echo "# building for $(OS)/$(ARCH)"
	@go build -o bin/$(firstword $(subst _, ,$*)) $(CMD_FOLDER)/$(firstword $(subst _, ,$*))/$(MAIN_FILE)