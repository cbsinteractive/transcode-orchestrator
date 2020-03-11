module github.com/cbsinteractive/video-transcoding-api

replace github.com/bitmovin/bitmovin-api-sdk-go => github.com/zsiec/bitmovin-api-sdk-go v1.30.0-alpha.0.0.20191206023358-8ff55f235fcf

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1 // indirect
	github.com/NYTimes/gizmo v1.3.5
	github.com/NYTimes/gziphandler v1.1.1
	github.com/aws/aws-sdk-go-v2 v0.19.0
	github.com/aws/aws-xray-sdk-go v0.9.4
	github.com/bitmovin/bitmovin-api-sdk-go v1.35.0-alpha.0
	github.com/cbsinteractive/hybrik-sdk-go v0.0.0-20191031180025-00f04ed90532
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/fsouza/ctxlogger v1.5.9
	github.com/fsouza/gizmo-stackdriver-logging v1.3.2
	github.com/getsentry/sentry-go v0.5.1
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/google/go-cmp v0.4.0
	github.com/google/gops v0.3.7
	github.com/gorilla/handlers v1.4.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kr/pretty v0.2.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
)

go 1.13
