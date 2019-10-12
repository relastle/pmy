.PHONY: main
main:
	# Remove comments from zsh source code
	sed '/^[[:blank:]]*#/d;s/#.*//' ./shell/pmy.zsh > ./_shell/_pmy.zsh
	# Make statik files
	statik -src=./_shell
	# build
	go build

.PHONY: clean
clean:
	rm -f ./anypm

.PHONY: bench
bench:
	go test -run=XXX -bench=.

.PHONY: lint
lint:
	golint ./main.go

.PHONY: docker
docker:
	docker build -t relastle/pmy:0.1.0 -f docker/Dockerfile .

.PHONY: integration_test
integration_test:
	(cd ./integration_test && go test -run .)
