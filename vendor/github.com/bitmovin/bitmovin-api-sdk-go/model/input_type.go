package model
type InputType string

// List of InputType
const (
	InputType_AKAMAI_NETSTORAGE InputType = "AKAMAI_NETSTORAGE"
	InputType_ASPERA InputType = "ASPERA"
	InputType_AZURE InputType = "AZURE"
	InputType_REDUNDANT_RTMP InputType = "REDUNDANT_RTMP"
	InputType_FTP InputType = "FTP"
	InputType_GENERIC_S3 InputType = "GENERIC_S3"
	InputType_GCS InputType = "GCS"
	InputType_HTTP InputType = "HTTP"
	InputType_HTTPS InputType = "HTTPS"
	InputType_LOCAL InputType = "LOCAL"
	InputType_RTMP InputType = "RTMP"
	InputType_S3 InputType = "S3"
	InputType_S3_ROLE_BASED InputType = "S3_ROLE_BASED"
	InputType_SFTP InputType = "SFTP"
	InputType_TCP InputType = "TCP"
	InputType_UDP InputType = "UDP"
	InputType_UDP_MULTICAST InputType = "UDP_MULTICAST"
	InputType_ZIXI InputType = "ZIXI"
	InputType_SRT InputType = "SRT"
)