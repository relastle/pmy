.PHONY: main
main: script
	go mod tidy
	go build
	go install

.PHONY: script
script: ./shell/pmy.zsh
	# Remove comments from zsh source code
	sed '/^[[:blank:]]*#/d;s/#.*//' ./shell/pmy.zsh > ./_shell/_pmy.zsh

.PHONY: bench
bench:
	go test -run=XXX -bench=.

.PHONY: docker
docker:
	docker build -t relastle/pmy:0.1.0 -f docker/Dockerfile .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test: lint
	go test -v ./src

.PHONY: integration_test
integration_test:
	(cd ./integration_test && go test -v -run .)

.PHONY: test_all
test_all:
	$(MAKE) main
	$(MAKE) test
	$(MAKE) integration_test


