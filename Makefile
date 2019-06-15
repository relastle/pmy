.PHONY: main
main:
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
