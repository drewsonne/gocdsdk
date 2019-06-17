.DEFAULT: test
SHELL:=/bin/bash
TEST?=$$(go list ./... |grep -v 'vendor')

GO_TARGETS= ./cli ./gocd ./gocd-*generator
GOCD_VERSION?= v18.12.0

format:
	gofmt -w -s .
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format
	$(MAKE) -C ./gocd-cli-action-generator/ format

lint:
	diff -u <(echo -n) <(gofmt -d -s main.go $(GO_TARGETS))
	golint -set_exit_status .

vet:
	go get ./...
	go vet ./...

test: vet lint
	go test -v -coverprofile=coverage.out -covermode=atomic ./...

before_install:
	@go get -u golang.org/x/lint/golint

build: deploy_on_develop

deploy_on_tag:
	git clean -df
	go get
	goreleaser --debug

deploy_on_develop:
	git clean -df
	go get
	goreleaser --debug --rm-dist --snapshot

testacc: provision-test-gocd
	bash scripts/wait-for-test-server.sh
	GOCD_ACC=1 $(MAKE) test

provision-test-gocd:
	cp godata/default.gocd.config.xml godata/server/config/cruise-config.xml
	docker rm -f gocd-server-test || true
	docker build -t gocd-server --build-arg UID=$(shell id -u) --build-arg GOCD_VERSION=${GOCD_VERSION} .
	docker run -p 8153:8153 -p 8154:8154 -d --name gocd-server-test gocd-server
