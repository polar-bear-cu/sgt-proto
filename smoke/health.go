package smoke

import healthv1 "github.com/polar-bear-cu/sgt-proto/gen/go/health/v1"

var (
	_ healthv1.HealthServiceClient
	_ healthv1.HealthServiceServer
	_ = healthv1.PingRequest{}
	_ = healthv1.PingResponse{}
)
