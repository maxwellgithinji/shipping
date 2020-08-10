module github.com/maxwellgithinji/shipping/shipping-cli-consignment

go 1.14

// replace github.com/maxwellgithinji/shipping/shipping-service-consignment => ../shipping-service-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/maxwellgithinji/shipping/shipping-service-consignment v0.0.0-20200808140052-44010d961e83
	github.com/micro/go-micro/v2 v2.9.1
)
