
setup:
#	apk add musl-dev go

build: setup
	@GOROOT=/usr/lib/go
	@GOPATH=/go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

all: build
	echo "All done"
