.PHONY: main
main: statik
	go mod tidy
	go build
	go install

statik: ./shell/pmy.zsh
	# Remove comments from zsh source code
	sed '/^[[:blank:]]*#/d;s/#.*//' ./shell/pmy.zsh > ./_shell/_pmy.zsh
	# Make statik files
	statik -src=./_shell

.PHONY: bench
bench:
	go test -run=XXX -bench=.

.PHONY: docker
docker:
	docker build -t relastle/pmy:0.1.0 -f docker/Dockerfile .

.PHONY: test
test: lint
	go test -v ./src

.PHONY: integration_test
integration_test: main test
	(cd ./integration_test && go test -v -run .)

.PHONY: lint
lint:
	golangci-lint run ./...

