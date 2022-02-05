WORKDIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GO111MODULE ?= auto
GOFLAGS ?= -mod=vendor

ifeq (, $(shell which richgo))
tst := go test
else
tst := richgo test
endif

ifneq (, $(VERBOSE))
tst := $(tst) -v
endif

.PHONY: test
test:
	$(tst) ./...

.PHONY: test-%
test-%:
	$(tst) ./$*


.PHONY: clean
clean:
	$(info Cleaning)
