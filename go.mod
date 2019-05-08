module github.com/hhy5861/gin-base

go 1.12

require (
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.3.1
	github.com/hhy5861/go-common v0.0.2
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829 // indirect
	github.com/urfave/cli v1.20.0
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	google.golang.org/grpc v1.20.1
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/hhy5861/go-common v0.0.2 => ../go-common
