APP ?= kolesa.scrapper

RELEASE ?= 1.0.0
COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

M = $(shell printf "\033[34;1m▶\033[0m")

all: vendor test build done

done:
	$(info $(M) done.)

.PHONY: vendor
vendor: prepare-dep ## Install dependicies
	$(info $(M) installing dependencies...)
	dep ensure

HAS_DEP := $(shell command -v dep;)

.PHONY: prepare-dep
prepare-dep: ## Install dep package manager
ifndef HAS_DEP
	$(info $(M) installing dep...)
	go get -u -v -d github.com/golang/dep/cmd/dep && \
	go install -v github.com/golang/dep/cmd/dep
endif
	
.PHONY: clean
clean:
	$(info $(M) cleaning build...)
	@rm -f bin/${APP}

.PHONY: build
build: clean ## Build program binary
	$(info $(M) building program...)
	go build -o bin/${APP}

.PHONY: run-old
run-old: ## Run in debug mode
	$(info $(M) running program...)
	bin/$(APP) -mode=old

.PHONY: run-new
run-new: ## Run in debug mode
	$(info $(M) running program...)
	bin/$(APP) -mode=new

.PHONY: run-all
run-all: ## Run in debug mode
	$(info $(M) running program...)
	bin/$(APP) -mode=all

.PHONY: test
test: ## Run test
	$(info $(M) running test...)
	go test -v -covermode=count -coverprofile=coverage.out


.PHONY: help
help: ## Show usage
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
