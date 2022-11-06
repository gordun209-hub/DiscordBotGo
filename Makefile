BINARY=main.out

build:
	go build -o $(BINARY) -v 

run:
	go build -o $(BINARY) main.go
	./$(BINARY)

clean:
	if [ -f $(BINARY) ] ; then rm $(BINARY) ; fi

test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

.PHONY: build run clean test fmt vet


