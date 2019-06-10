.PHONY: main
main:
	go build

.PHONY: clean
clean:
	rm -f ./anypm

.PHONY: bench
bench:
	go test -run=XXX -bench=.
