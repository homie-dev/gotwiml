BIN = $(GOPATH)/bin

goimports=$(BIN)/goimports
$(goimports):
	go get golang.org/x/tools/cmd/goimports

fmt:
	$(goimports) -w -local twiml .
