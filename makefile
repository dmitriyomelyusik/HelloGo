HELLOGO := $(HOME)/Go/projects/HelloWorld/src/hello
GOPATH := $(HOME)/go

.PHONY: run
run: test
	chromium-browser http://localhost:8080
	go run $(HELLOGO)/helloworld.go

.PHONY: test
test: lint
	go test $(HELLOGO)/helloworld_test.go


.PHONY: lint
lint:
	gometalinter $(HELLOGO)/helloworld.go
	gometalinter $(HELLOGO)/helloworld_test.go


