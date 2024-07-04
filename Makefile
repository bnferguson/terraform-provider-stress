VERSION = 1.0.0
OS_ARCHS = linux_amd64 darwin_arm64
TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=stress
DIR=~/.terraform.d/plugins
GO_CLIENT_VERSION=master

default: build

build: fmtcheck
	go install

install: fmtcheck
	mkdir -vp $(DIR)
	go build -o $(DIR)/terraform-provider-stress

uninstall:
	@rm -vf $(DIR)/terraform-provider-stress

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout=120s

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

build-all: fmtcheck
	mkdir -vp terraform.d/plugins/terraform.volcanicislandfortress.com/bnferguson/stress/1.0.0/linux_amd64
	GOOS=linux GOARCH=amd64 go build -o terraform.d/plugins/terraform.volcanicislandfortress.com/bnferguson/stress/1.0.0/linux_amd64/terraform-provider-stress .
	mkdir -vp terraform.d/plugins/terraform.volcanicislandfortress.com/bnferguson/stress/1.0.0/darwin_arm64
	GOOS=darwin GOARCH=arm64 go build -o terraform.d/plugins/terraform.volcanicislandfortress.com/bnferguson/stress/1.0.0/darwin_arm64/terraform-provider-stress .

build-and-package: build-all
	@for os_arch in $(OS_ARCHS); do \
		echo '{ "version": "$(VERSION)", "architectures": ["amd64"], "os": ["linux", "darwin"], "filename": "terraform-provider-stress", "protocols": ["5.0"] }' > terraform-provider-stress_$(VERSION)_manifest.json ; \
		zip -j terraform-provider-stress_$(VERSION)_$$os_arch.zip terraform.d/plugins/terraform.volcanicislandfortress.com/bnferguson/stress/$(VERSION)/$$os_arch/terraform-provider-stress terraform-provider-stress_$(VERSION)_manifest.json ; \
	done

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)


.PHONY: build build-all build-and-package test testacc vet fmt fmtcheck errcheck test-compile
