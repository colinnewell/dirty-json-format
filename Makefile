all: djf

djf: cmd/*.go
	go build -o djf cmd/*.go

clean:
	rm djf

install:
	cp djf /usr/local/bin

lint:
	golangci-lint run
	golint ./...
