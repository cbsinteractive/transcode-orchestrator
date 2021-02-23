package job


const old0 = `{
	"providers": ["mediaconvert"],
	"preset": {
		"name": "whatever",
		"container": "mxf",
		"rateControl": "CBR",
		"video": {"height": "1080","width": "1920","codec": "xdcam","profile": "hd422","bitrate": "5000000","gopSize": "60","gopMode": "fixed","interlaceMode": "interlaced"},
		"audio": {"codec": "pcm","discreteTracks": true}
	}
}`
const old1 = `{
	"provider": "mediaconvert",
	"source": "s3://vtg-as-test-bucket/mxf/test/in.mp4",
	"destinationBasePath": "s3://vtg-as-test-bucket/mxf/test",
	"outputs": [
		{
			"preset": "whatever",
			"fileName": "out.mxf"
		}
	]
}`

const new0 = `{
	"jobID": "0000001",
	"provider": "mediaconvert",
	"source": "s3://vtg-as-test-bucket/mxf/test/in.mp4",
	"destinationBasePath": "s3://vtg-as-test-bucket/mxf/test",
	"outputs": [{
			"preset": {
				"name": "whatever",
				"container": "mxf",
				"rateControl": "CBR",
				"video": {"height": "1080","width": "1920","codec": "xdcam","profile": "hd422","bitrate": "5000000","gopSize": "60","gopMode": "fixed","interlaceMode": "interlaced"},
				"audio": {"codec": "pcm","discreteTracks": true}
			},
			"fileName": "out.mxf"
		}
	]
}`
