package model
import (
	"time"
)

type AwsAccountRegionSettings struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Limit for the amount of running encodings at a time. Leave empty for no limit.
	LimitParallelEncodings *int64 `json:"limitParallelEncodings,omitempty"`
	// Maximum amount of encoding coordinators and workers allowed in this region at any time. Leave empty for no limit.
	MaximumAmountOfCoordinatorsAndWorkersInRegion *int64 `json:"maximumAmountOfCoordinatorsAndWorkersInRegion,omitempty"`
	// Limit the amount of money to spend in this region on this account. Leave empty for no limit.
	MaxMoneyToSpendPerMonth *float64 `json:"maxMoneyToSpendPerMonth,omitempty"`
	// Id of the security group for encoding instances (required)
	SecurityGroupId string `json:"securityGroupId,omitempty"`
	// Id of the subnet for encoding instances (required)
	SubnetId string `json:"subnetId,omitempty"`
	// Which machine types are allowed to be deployed. Leave empty for no machine type restrictions.
	MachineTypes []string `json:"machineTypes,omitempty"`
	// Custom SSH port. Valid values: 1 - 65535. Leave empty if the default SSH port 22 is OK.
	SshPort *int32 `json:"sshPort,omitempty"`
}

