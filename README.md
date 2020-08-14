<p align="center">
  <img width="450" src="https://p191.p3.n0.cdn.getcloudapp.com/items/wbu0k6m1/path1059.png?v=44e064cde05df48d7a8ce13ca62e760c">
</p>
---

[![Go Report Card](https://goreportcard.com/badge/github.com/cbsinteractive/transcode-orchestrator)](https://goreportcard.com/report/github.com/cbsinteractive/transcode-orchestrator)

The transcode orchestrator provides an agnostic API to transcode media assets
across different cloud services. Currently, it supports the following
providers:

- [Bitmovin](http://bitmovin.com)
- [Hybrik](https://www.hybrik.com)
- [MediaConvert](https://aws.amazon.com/mediaconvert)
- [Flock (CBS internal provider)](https://github.com/cbsinteractive/flock)

## Setting Up

With [latest Go](https://golang.org/dl/) installed, make sure to export the follow
environment variables:

### Providers configuration

#### For [Bitmovin](http://bitmovin.com)

```
export BITMOVIN_API_KEY=your.api.key
export BITMOVIN_AWS_ACCESS_KEY_ID=your.access.key.id
export BITMOVIN_AWS_SECRET_ACCESS_KEY=your.secret.access.key
export BITMOVIN_AWS_STORAGE_REGION=your.s3.region.such.as.US_EAST_1.or.EU_WEST_1
export BITMOVIN_GCS_ACCESS_KEY_ID=your.gcs.access.key.id
export BITMOVIN_GCS_SECRET_ACCESS_KEY=your.gcs.secret.access.key
export BITMOVIN_GCS_STORAGE_REGION=your.s3.region.such.as.US_EAST_1.or.EU_WEST_1
export BITMOVIN_DESTINATION=s3://your-s3-bucket
export BITMOVIN_ENCODING_REGION=your.provider.region.such.as.AWS_US_EAST_1.or.GOOGLE_EUROPE_WEST_1
export BITMOVIN_ENCODING_VERSION=STABLE.or.BETA
```

#### For [Hybrik](https://www.hybrik.com)

```
export HYBRIK_URL=your.hybrik.api.endpoint.such.as.https://api_demo.hybrik.com/v1
export HYBRIK_COMPLIANCE_DATE=20170601
export HYBRIK_OAPI_KEY=your.hybrik.oapi.key
export HYBRIK_OAPI_SECRET=your.hybrik.oapi.secret
export HYBRIK_AUTH_KEY=your.hybrik.auth.key
export HYBRIK_AUTH_SECRET=your.hybrik.auth.secret
export HYBRIK_GCP_CREDENTIALS_KEY=your.hybrik.gcp.credentials.key
export HYBRIK_DESTINATION=s3://your-s3-bucket
export HYBRIK_PRESET_PATH=video-transcoding-api-presets
```

``HYBRIK_PRESET_PATH`` is optional and defines the folder presets will be
stored in. If not specified, it will default to
'video-transcoding-api-presets'.

#### For [MediaConvert](https://aws.amazon.com/mediaconvert/)

```
export MEDIACONVERT_AWS_ACCESS_KEY_ID=your.access.key.id
export MEDIACONVERT_AWS_SECRET_ACCESS_KEY=your.secret.access.key
export MEDIACONVERT_AWS_REGION="us-east-1"
export MEDIACONVERT_ENDPOINT=your.mediaconvert.endpoint
export MEDIACONVERT_QUEUE_ARN=your.queue.arn
export MEDIACONVERT_PREFERRED_QUEUE_ARN=your.preferred.queue.arn
export MEDIACONVERT_ROLE_ARN=your.iam.role.arn
export MEDIACONVERT_DESTINATION=s3://your-s3-bucket
```

#### For [Flock](https://github.com/cbsinteractive/flock)

```
export FLOCK_ENDPOINT=your.flock.endpoint
export FLOCK_CREDENTIAL=your.flock.auth.secret
```

### Database configuration

In order to store preset maps and job statuses we need a Redis instance
running. Learn how to set up and run a Redis server
[here](http://redis.io/topics/quickstart). With the Redis instance running, set
its configuration variables:

```
export REDIS_ADDR=192.0.2.31
export REDIS_PASSWORD=p4ssw0rd.here
```

If you are running Redis in the same host of the API and on the default port
(6379) the API will automatically find the instance and connect to it.

With all environment variables set and redis up and running, clone this
repository and run:

```
$ git clone https://github.com/cbsinteractive/transcode-orchestrator.git
$ make run
```

## Running tests

```
$ make test
```

## License

- This code is under [Apache 2.0
  license](https://github.com/cbsinteractive/transcode-orchestrator/blob/master/LICENSE).
