package model

type SpekeDrmProvider struct {
	// URL of the endpoint (required)
	Url string `json:"url,omitempty"`
	// Your username for Basic Authentication
	Username string `json:"username,omitempty"`
	// Your password for Basic Authentication
	Password string `json:"password,omitempty"`
	// AWS role that will be assumed for the key exchange in case the provider runs on AWS.  During the key exchange the role will be assumed to be able to access the key provider.  This role has to be created by the owner of the account with the SPEKE server. For Bitmovin to be able to assume this role, the following has to be added to the trust policy of the role:  ``` {   \"Effect\": \"Allow\",   \"Principal\": {     \"AWS\": \"arn:aws:iam::630681592166:user/bitmovinCustomerSpekeAccess\"   },   \"Action\": \"sts:AssumeRole\",   \"Condition\": {     \"StringEquals\": {       \"sts:ExternalId\": \"{{externalId}}\"     }   } } ``` It is recommended to also set the {{externalId}} due to security reasons but it can also be ommitted.  Additionally the role needs a policy similar to the following to be able to invoke the API gateway: ``` {   \"Version\": \"2012-10-17\",   \"Statement\": [     {       \"Effect\": \"Allow\",       \"Action\": [         \"execute-api:Invoke\"       ],       \"Resource\": [         \"arn:aws:execute-api:{{region}}:*:*_/_*_/POST/_*\"       ]     }   ] } ``` where `{{region}}` is the region of the API gateway (for example `us-west-2`), the same has to be set in the property 'gatewayRegion'. It's also possible to set `{{region}` to `*` to give the role access to all regions. 
	RoleArn string `json:"roleArn,omitempty"`
	// External ID used together with the IAM role identified by `roleArn` to assume access to the SPEKE server on AWS. 
	ExternalId string `json:"externalId,omitempty"`
	// Describes the region of the AWS API Gateway that is used to access the SPEKE server. This property is mandatory when setting 'roleArn' and has to indicate in which region the AWS API Gateway is setup. This usually corresponds to the `{{region}}` one sets in the execute-api policy for the role as described in 'roleArn'. 
	GatewayRegion string `json:"gatewayRegion,omitempty"`
}

