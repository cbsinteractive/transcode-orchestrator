package provider

type Cloud = string

const (
	CloudAWS Cloud = "aws"
	CloudGCP Cloud = "gcp"
)

type GCPCloudRegion = string

const (
	GCPRegionUSEast1    GCPCloudRegion = "us-east1"
	GCPRegionUSEast4    GCPCloudRegion = "us-east4"
	GCPRegionUSWest1    GCPCloudRegion = "us-west1"
	GCPRegionUSWest2    GCPCloudRegion = "us-west2"
	GCPRegionUSCentral1 GCPCloudRegion = "us-central1"
)

type AWSCloudRegion = string

const (
	AWSRegionUSEast1 AWSCloudRegion = "us-east-1"
	AWSRegionUSEast2 AWSCloudRegion = "us-east-2"
	AWSRegionUSWest1 AWSCloudRegion = "us-west-1"
	AWSRegionUSWest2 AWSCloudRegion = "us-west-2"
)
