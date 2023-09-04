<<<<<<< HEAD
TEST?=$$(go list ./... | grep -v github.com/hashicorp/terraform-provider-google/scripts)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=google
DIR_NAME=google

default: build

build: lint generate
	go install

test: lint generate
	go test $(TESTARGS) -timeout=30s $(TEST)

testacc: generate
	TF_ACC=1 TF_SCHEMA_PANIC_ON_ERROR=1 go test $(TEST) -v $(TESTARGS) -timeout 240m -ldflags="-X=github.com/hashicorp/terraform-provider-google/version.ProviderVersion=acc"
=======
TEST?=$$(go list ./... | grep -v github.com/hashicorp/terraform-provider-google-beta/scripts)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=google
DIR_NAME=google-beta

default: build

build: lint
	go install

test: lint
	go test $(TESTARGS) -timeout=30s $(TEST)

testacc: lint
	TF_ACC=1 TF_SCHEMA_PANIC_ON_ERROR=1 go test $(TEST) -v $(TESTARGS) -timeout 240m -ldflags="-X=github.com/hashicorp/terraform-provider-google-beta/version.ProviderVersion=acc"
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w -s ./$(DIR_NAME)

# Currently required by tf-deploy compile
fmtcheck:
	sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

vet:
	go vet

lint: fmtcheck vet
<<<<<<< HEAD

generate:
	go generate  ./...
=======
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(DIR_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

docscheck:
	@sh -c "'$(CURDIR)/scripts/docscheck.sh'"
<<<<<<< HEAD
=======

.PHONY: build test testacc fmt fmtcheck vet lint  errcheck test-compile website website-test docscheck
>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35

.PHONY: build test testacc fmt fmtcheck vet lint test-compile website website-test docscheck generate
