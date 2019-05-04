all: media

clean:
	@rm -rf ./gin-base-linux-amd64*

media:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gin-base-linux-amd64 -v .