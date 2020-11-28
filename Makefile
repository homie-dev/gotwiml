BIN = $(GOPATH)/bin

goimports=$(BIN)/goimports
$(goimports):
	go get golang.org/x/tools/cmd/goimports

fmt:
	$(goimports) -w -local github.com/homie-dev/gotwiml .

.PHONY: test
test:
	go test -p 20 $(shell go list ./...)
