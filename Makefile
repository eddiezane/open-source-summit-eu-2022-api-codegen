# https://marcofranssen.nl/manage-go-tools-via-go-modules

.PHONY: download
download:
	@echo Download go.mod dependencies
	@go mod download

.PHONY: install-tools
install-tools: download
	@echo Installing tools from tools.go
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: generate
generate:
	./hack/update-codegen.sh
