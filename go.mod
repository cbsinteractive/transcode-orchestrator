module github.com/cbsinteractive/transcode-orchestrator

replace github.com/bitmovin/bitmovin-api-sdk-go => github.com/zsiec/bitmovin-api-sdk-go v1.30.0-alpha.0.0.20191206023358-8ff55f235fcf

replace github.com/cbsinteractive/transcode-orchestrator/client => ./client

require (
	github.com/NYTimes/gizmo v1.3.5 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/aws/aws-sdk-go v1.30.9
	github.com/aws/aws-sdk-go-v2 v0.24.0
	github.com/awslabs/smithy-go v0.0.0-20200421200441-f1e89484c1b9 // indirect
	github.com/bitmovin/bitmovin-api-sdk-go v1.37.0-alpha.0
	github.com/cbsinteractive/hybrik-sdk-go v0.0.0-20191031180025-00f04ed90532
	github.com/cbsinteractive/pkg/timecode v0.0.0-20200805005557-c8cd96fa0686
	github.com/cbsinteractive/pkg/video v0.0.4 // indirect
	github.com/cbsinteractive/transcode-orchestrator/client v0.0.0-00010101000000-000000000000
	github.com/fsouza/ctxlogger v1.5.9 // indirect
	github.com/fsouza/gizmo-stackdriver-logging v1.3.2 // indirect
	github.com/getsentry/sentry-go v0.6.0
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/google/go-cmp v0.4.0
	github.com/google/gops v0.3.7 // indirect
	github.com/gorilla/handlers v1.4.2 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.5.0
	github.com/zsiec/pkg/tracing v0.0.0-20200316013157-874eb6019248
	github.com/zsiec/pkg/xrayutil v0.0.0-20200316013157-874eb6019248 // indirect
)

go 1.14
