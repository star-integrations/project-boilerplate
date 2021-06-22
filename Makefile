API_GEN_VERSION := v2.1.0
FIRESTORE_REPO_VERSION := v1.0.0
GOLANGCI_LINT_VERSION := 1.41.0
SWAG_VERSION := 1.7.0
MOCKGEN_VERSION := 1.6.0

OS_NAME := `echo $(shell uname -s) | tr A-Z a-z`
MACHINE_TYPE := $(shell uname -m)

.PHONY: init
init: bootstrap
	test -f .env || cp .env.template .env

.PHONY: bootstrap
bootstrap: bootstrap_api_gen bootstrap_firestore_repo bootstrap_golangci_lint bootstrap_swag bootstrap_mockgen

.PHONY: bootstrap_api_gen
bootstrap_api_gen:
	mkdir -p ./bin
	curl -s -L -o ./bin/api_gen.tar.gz https://github.com/go-generalize/api_gen/releases/download/$(API_GEN_VERSION)/api_gen_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && \
	tar xzf api_gen.tar.gz && \
	rm *.tar.gz

.PHONY: server_generate
server_generate:
	cd ./back && ../bin/api_gen server -p server -o server ./interfaces

.PHONY: client_generate
client_generate:
	mkdir -p ./front/lib/api
	cd ./front/lib/api && ../../../bin/api_gen client ts ../../../back/interfaces

.PHONY: generate
generate: server_generate client_generate go_generate

.PHONY: bootstrap_firestore_repo
bootstrap_firestore_repo:
	mkdir -p bin
	curl -L -o ./bin/firestore-repo.tar.gz https://github.com/go-generalize/firestore-repo/releases/download/$(FIRESTORE_REPO_VERSION)/firestore-repo_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && tar xzf firestore-repo.tar.gz && rm *.tar.gz

.PHONY: bootstrap_golangci_lint
bootstrap_golangci_lint:
	mkdir -p bin
	curl -s -L -o ./bin/golangci-lint.tar.gz https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/golangci-lint-$(GOLANGCI_LINT_VERSION)-$(OS_NAME)-amd64.tar.gz
	cd ./bin && \
	tar xzf golangci-lint.tar.gz && \
	mv golangci-lint-$(GOLANGCI_LINT_VERSION)-$(shell uname -s)-amd64/golangci-lint golangci-lint && \
	rm -rf golangci-lint-$(GOLANGCI_LINT_VERSION)-$(shell uname -s)-amd64 *.tar.gz

.PHONY: lint
lint:
	./bin/golangci-lint run --config=".github/.golangci.yml" --fast ./...

.PHONY: bootstrap_swag
bootstrap_swag:
	mkdir -p ./bin
	curl -s -L -o ./bin/swag.tar.gz https://github.com/swaggo/swag/releases/download/v$(SWAG_VERSION)/swag_$(SWAG_VERSION)_$(OS_NAME)_$(MACHINE_TYPE).tar.gz
	cd ./bin && \
	tar xzf swag.tar.gz && \
	rm *.tar.gz

.PHONY: swag
swag:
	cd back && make swag

.PHONY: bootstrap_mockgen
bootstrap_mockgen:
	mkdir -p ./bin
	curl -s -L -o ./bin/mockgen.tar.gz https://github.com/golang/mock/releases/download/v$(MOCKGEN_VERSION)/mock_$(MOCKGEN_VERSION)_$(OS_NAME)_amd64.tar.gz
	cd ./bin && \
	tar xzf mockgen.tar.gz && \
	mv mock_$(MOCKGEN_VERSION)_$(OS_NAME)_amd64/mockgen ./mockgen && \
	rm -rf mock_$(MOCKGEN_VERSION)_$(OS_NAME)_amd64 *.tar.gz

.PHONY: go_generate
go_generate:
	go generate ./...

.PHONY: run_server
run_server:
	cd back/cmd && go run .

.PHONY: test
test:
	go test -v ./... | grep -v 'no test files'

.PHONY: docker_cleanup
docker_cleanup:
	docker ps -a | awk 'NR>1 {print $1}' | xargs docker rm -f
	docker images | awk 'NR>1 {print $3}' | xargs docker rmi -f

.PHONY: prune
prune:
	git remote prune origin
	git fetch origin --prune 'refs/tags/*:refs/tags/*'
	git branch --merged | egrep -v '^ *(develop|main|\*)' | xargs git branch -d
