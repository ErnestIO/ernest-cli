VERSION=`cat ./VERSION`

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

install:
	go install -v ${LDFLAGS}

build:
	go build -o ernest -v ${LDFLAGS}

test:
	go test -v ./...

cover:
	go test -coverprofile cover.out

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

dev-deps: deps
	go get github.com/alecthomas/gometalinter
	gometalinter --install

lint:
	gometalinter --config .linter.conf

dist: dist-linux dist-darwin dist-windows

dist-linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ernest-${VERSION}-linux-amd64
	zip ernest-${VERSION}-linux-amd64.zip ernest-${VERSION}-linux-amd64 README.md LICENSE
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o ernest-${VERSION}-linux-386
	zip ernest-${VERSION}-linux-386.zip ernest-${VERSION}-linux-386 README.md LICENSE

dist-darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ernest-${VERSION}-darwin-amd64
	zip ernest-${VERSION}-darwin-amd64.zip ernest-${VERSION}-darwin-amd64 README.md LICENSE
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o ernest-${VERSION}-darwin-386
	zip ernest-${VERSION}-darwin-386.zip ernest-${VERSION}-darwin-386 README.md LICENSE

dist-windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ernest-${VERSION}-windows-amd64.exe
	zip ernest-${VERSION}-windows-amd64.zip ernest-${VERSION}-windows-amd64.exe README.md LICENSE
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ernest-${VERSION}-windows-386.exe
	zip ernest-${VERSION}-windows-386.zip ernest-${VERSION}-windows-386.exe README.md LICENSE

assets:
	cd helper && go-bindata -pkg helper -nocompress lang

clean:
	go clean
	rm -rf ernest-*
