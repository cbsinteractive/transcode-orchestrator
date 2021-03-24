module github.com/cbsinteractive/transcode-orchestrator

replace github.com/bitmovin/bitmovin-api-sdk-go => github.com/zsiec/bitmovin-api-sdk-go v1.30.0-alpha.0.0.20191206023358-8ff55f235fcf

require (
	github.com/aws/aws-sdk-go v1.30.9
	github.com/aws/aws-sdk-go-v2 v0.24.0
	github.com/bitmovin/bitmovin-api-sdk-go v1.37.0-alpha.0
	github.com/cbsinteractive/hybrik-sdk-go v0.0.0-20191031180025-00f04ed90532
	github.com/cbsinteractive/pkg/timecode v0.0.0-20200805005557-c8cd96fa0686
	github.com/cbsinteractive/pkg/video v0.0.4
	github.com/getsentry/sentry-go v0.6.0
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.5.0
	github.com/zsiec/pkg/tracing v0.0.0-20200316013157-874eb6019248
	golang.org/x/text v0.3.2 // indirect
)

go 1.14
