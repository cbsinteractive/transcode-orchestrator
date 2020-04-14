package model

type PlayerThirdPartyLicensing struct {
	// URL to your license check server (required)
	LicenseCheckServer string `json:"licenseCheckServer,omitempty"`
	// Timeout in ms (required)
	LicenseCheckTimeout *int32 `json:"licenseCheckTimeout,omitempty"`
	// Specify if the Licensing Request should fail or not on Third Party Licensing Error (required)
	ErrorAction PlayerThirdPartyLicensingErrorAction `json:"errorAction,omitempty"`
	// Specify if the Licensing Request should fail or not on Third Party Licensing timeout (required)
	TimeoutAction PlayerThirdPartyLicensingErrorAction `json:"timeoutAction,omitempty"`
}

