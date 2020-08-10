module github.com/maxwellgithinji/shipping/shipping-cli-consignment

go 1.14

// replace github.com/maxwellgithinji/shipping/shipping-service-consignment => ../shipping-service-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/maxwellgithinji/shipping/shipping-service-consignment v0.0.0-20200810124425-5e6c172d8c20
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/sys v0.0.0-20200810151505-1b9f1253b3ed // indirect
	golang.org/x/tools v0.0.0-20200810190217-c1903db4dbfe // indirect
)
