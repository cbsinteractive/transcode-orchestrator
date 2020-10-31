package bitmovin

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
)

var awsCloudRegions = map[model.AwsCloudRegion]struct{}{
	model.AwsCloudRegion_US_EAST_1: {}, model.AwsCloudRegion_US_EAST_2: {}, model.AwsCloudRegion_US_WEST_1: {},
	model.AwsCloudRegion_US_WEST_2: {}, model.AwsCloudRegion_EU_WEST_1: {}, model.AwsCloudRegion_EU_CENTRAL_1: {},
	model.AwsCloudRegion_AP_SOUTHEAST_1: {}, model.AwsCloudRegion_AP_SOUTHEAST_2: {}, model.AwsCloudRegion_AP_NORTHEAST_1: {},
	model.AwsCloudRegion_AP_NORTHEAST_2: {}, model.AwsCloudRegion_AP_SOUTH_1: {}, model.AwsCloudRegion_SA_EAST_1: {},
	model.AwsCloudRegion_EU_WEST_2: {}, model.AwsCloudRegion_EU_WEST_3: {}, model.AwsCloudRegion_CA_CENTRAL_1: {},
}

var regionByCloud = map[string]map[string]model.CloudRegion{
	provider.CloudAWS: {
		provider.AWSRegionUSEast1: model.CloudRegion_AWS_US_EAST_1,
		provider.AWSRegionUSEast2: model.CloudRegion_AWS_US_EAST_2,
		provider.AWSRegionUSWest1: model.CloudRegion_AWS_US_WEST_1,
		provider.AWSRegionUSWest2: model.CloudRegion_AWS_US_WEST_2,
	},
	provider.CloudGCP: {
		provider.GCPRegionUSEast1:    model.CloudRegion_GOOGLE_US_EAST_1,
		provider.GCPRegionUSEast4:    model.CloudRegion_GOOGLE_US_EAST_4,
		provider.GCPRegionUSWest1:    model.CloudRegion_GOOGLE_US_WEST_1,
		provider.GCPRegionUSWest2:    model.CloudRegion_GOOGLE_US_WEST_2,
		provider.GCPRegionUSCentral1: model.CloudRegion_GOOGLE_US_CENTRAL_1,
	},
}

var cloudRegions = map[model.CloudRegion]struct{}{
	model.CloudRegion_AWS_US_EAST_1:                   {},
	model.CloudRegion_AWS_US_EAST_2:                   {},
	model.CloudRegion_AWS_US_WEST_1:                   {},
	model.CloudRegion_AWS_US_WEST_2:                   {},
	model.CloudRegion_AWS_EU_WEST_1:                   {},
	model.CloudRegion_AWS_EU_CENTRAL_1:                {},
	model.CloudRegion_AWS_AP_SOUTHEAST_1:              {},
	model.CloudRegion_AWS_AP_SOUTHEAST_2:              {},
	model.CloudRegion_AWS_AP_NORTHEAST_1:              {},
	model.CloudRegion_AWS_AP_NORTHEAST_2:              {},
	model.CloudRegion_AWS_AP_SOUTH_1:                  {},
	model.CloudRegion_AWS_SA_EAST_1:                   {},
	model.CloudRegion_AWS_EU_WEST_2:                   {},
	model.CloudRegion_AWS_EU_WEST_3:                   {},
	model.CloudRegion_AWS_CA_CENTRAL_1:                {},
	model.CloudRegion_GOOGLE_US_CENTRAL_1:             {},
	model.CloudRegion_GOOGLE_US_EAST_1:                {},
	model.CloudRegion_GOOGLE_ASIA_EAST_1:              {},
	model.CloudRegion_GOOGLE_EUROPE_WEST_1:            {},
	model.CloudRegion_GOOGLE_US_WEST_1:                {},
	model.CloudRegion_GOOGLE_ASIA_EAST_2:              {},
	model.CloudRegion_GOOGLE_ASIA_NORTHEAST_1:         {},
	model.CloudRegion_GOOGLE_ASIA_SOUTH_1:             {},
	model.CloudRegion_GOOGLE_ASIA_SOUTHEAST_1:         {},
	model.CloudRegion_GOOGLE_AUSTRALIA_SOUTHEAST_1:    {},
	model.CloudRegion_GOOGLE_EUROPE_NORTH_1:           {},
	model.CloudRegion_GOOGLE_EUROPE_WEST_2:            {},
	model.CloudRegion_GOOGLE_EUROPE_WEST_4:            {},
	model.CloudRegion_GOOGLE_NORTHAMERICA_NORTHEAST_1: {},
	model.CloudRegion_GOOGLE_SOUTHAMERICA_EAST_1:      {},
	model.CloudRegion_GOOGLE_US_EAST_4:                {},
	model.CloudRegion_GOOGLE_US_WEST_2:                {},
	model.CloudRegion_AZURE_EUROPE_WEST:               {},
	model.CloudRegion_AZURE_US_WEST2:                  {},
	model.CloudRegion_AZURE_US_EAST:                   {},
	model.CloudRegion_AZURE_AUSTRALIA_SOUTHEAST:       {},
	model.CloudRegion_NORTH_AMERICA:                   {},
	model.CloudRegion_SOUTH_AMERICA:                   {},
	model.CloudRegion_EUROPE:                          {},
	model.CloudRegion_AFRICA:                          {},
	model.CloudRegion_ASIA:                            {},
	model.CloudRegion_AUSTRALIA:                       {},
	model.CloudRegion_AWS:                             {},
	model.CloudRegion_GOOGLE:                          {},
	model.CloudRegion_KUBERNETES:                      {},
	model.CloudRegion_EXTERNAL:                        {},
	model.CloudRegion_AUTO:                            {},
}
