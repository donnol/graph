.PHONY:

bench:
	go test -benchmem -bench=^Bench -run=^$$ . -v
