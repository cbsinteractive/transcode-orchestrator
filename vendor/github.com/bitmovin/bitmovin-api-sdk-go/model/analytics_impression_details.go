package model

type AnalyticsImpressionDetails struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Is an ad playing. 0 indicates no, 1 indicates yes
	Ad *int64 `json:"ad,omitempty"`
	// Collector version
	AnalyticsVersion string `json:"analyticsVersion,omitempty"`
	// Autonomous System Number inferred from the IP address
	Asn *int64 `json:"asn,omitempty"`
	// Audio Bitrate
	AudioBitrate *int64 `json:"audioBitrate,omitempty"`
	// Selected audio language
	AudioLanguage string `json:"audioLanguage,omitempty"`
	// Autoplay enabled
	Autoplay *bool `json:"autoplay,omitempty"`
	// Browser name
	Browser string `json:"browser,omitempty"`
	// Browser version major
	BrowserVersionMajor string `json:"browserVersionMajor,omitempty"`
	// Browser version minor
	BrowserVersionMinor string `json:"browserVersionMinor,omitempty"`
	// Milliseconds the player buffered
	Buffered *int64 `json:"buffered,omitempty"`
	// CDN Provider
	CdnProvider string `json:"cdnProvider,omitempty"`
	// City
	City string `json:"city,omitempty"`
	// Current time of the client
	ClientTime *int64 `json:"clientTime,omitempty"`
	// Country
	Country string `json:"country,omitempty"`
	// Custom user ID
	CustomUserId string `json:"customUserId,omitempty"`
	// Free form data set via the customData1 field in the analytics collector configuration
	CustomData1 string `json:"customData1,omitempty"`
	// Free form data set via the customData2 field in the analytics collector configuration
	CustomData2 string `json:"customData2,omitempty"`
	// Free form data set via the customData3 field in the analytics collector configuration
	CustomData3 string `json:"customData3,omitempty"`
	// Free form data set via the customData4 field in the analytics collector configuration
	CustomData4 string `json:"customData4,omitempty"`
	// Free form data set via the customData5 field in the analytics collector configuration
	CustomData5 string `json:"customData5,omitempty"`
	// Type of the device detected via User Agent
	DeviceType string `json:"deviceType,omitempty"`
	// Domain the player was loaded on (.www is stripped away)
	Domain string `json:"domain,omitempty"`
	// Time in milliseconds it took the DRM server to respond
	DrmLoadTime *int64 `json:"drmLoadTime,omitempty"`
	// DRM system used for this impression
	DrmType string `json:"drmType,omitempty"`
	// Dropped frames during playback
	DroppedFrames *int64 `json:"droppedFrames,omitempty"`
	// Duration of the sample in milliseconds
	Duration *int64 `json:"duration,omitempty"`
	// Error code
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// Error message
	ErrorMessage string `json:"errorMessage,omitempty"`
	// A/B test experiment name
	ExperimentName string `json:"experimentName,omitempty"`
	// Random UUID that is used to identify a sessions (required)
	ImpressionId string `json:"impressionId,omitempty"`
	// IP Address of the client
	IpAddress string `json:"ipAddress,omitempty"`
	// Is chromecast active
	IsCasting *bool `json:"isCasting,omitempty"`
	// Is the stream live or VoD
	IsLive *bool `json:"isLive,omitempty"`
	// Is the player muted
	IsMuted *bool `json:"isMuted,omitempty"`
	// The users Internet Service Provider inferred via the IP address
	Isp string `json:"isp,omitempty"`
	// Language set in the browser
	Language string `json:"language,omitempty"`
	// Analytics license key
	LicenseKey string `json:"licenseKey,omitempty"`
	// Operating system
	OperatingSystem string `json:"operatingSystem,omitempty"`
	// Operating system version major
	OperatingSystemVersionMajor string `json:"operatingSystemVersionMajor,omitempty"`
	// Operating system version minor
	OperatingSystemVersionMinor string `json:"operatingSystemVersionMinor,omitempty"`
	// Time in milliseconds the page took to load
	PageLoadTime *int32 `json:"pageLoadTime,omitempty"`
	// Player load type. 1 = Foreground, 2 = Background
	PageLoadType *int32 `json:"pageLoadType,omitempty"`
	// path on the website
	Path string `json:"path,omitempty"`
	// Milliseconds the player was paused
	Paused *int64 `json:"paused,omitempty"`
	// Platform the player is running on (web, android, ios)
	Platform string `json:"platform,omitempty"`
	// Milliseconds the player played
	Played *int64 `json:"played,omitempty"`
	// Video player being used for this session
	Player string `json:"player,omitempty"`
	// Player license key
	PlayerKey string `json:"playerKey,omitempty"`
	// Time in milliseconds the player took to start up
	PlayerStartuptime *int32 `json:"playerStartuptime,omitempty"`
	// HTML or native playback
	PlayerTech string `json:"playerTech,omitempty"`
	// Player software version
	PlayerVersion string `json:"playerVersion,omitempty"`
	// Geographic region (ISO 3166-2 code)
	Region string `json:"region,omitempty"`
	// Screen as reported by the browser
	ScreenHeight *int32 `json:"screenHeight,omitempty"`
	// Screen as reported by the browser
	ScreenWidth *int32 `json:"screenWidth,omitempty"`
	// Milliseconds it took the player to seek
	Seeked *int64 `json:"seeked,omitempty"`
	// Sequence number of the sample in which it occured in the session
	SequenceNumber *int64 `json:"sequenceNumber,omitempty"`
	// Video size (FULLSCREEN or WINDOW)
	Size string `json:"size,omitempty"`
	// Combination of player- and videoStartuptime
	StartupTime *int32 `json:"startupTime,omitempty"`
	// Internal state of the analytics state machine
	State string `json:"state,omitempty"`
	// Format of the stream (HLS, DASH, Progressive MP4)
	StreamFormat string `json:"streamFormat,omitempty"`
	// Subtitle enabled
	SubtitleEnabled *bool `json:"subtitleEnabled,omitempty"`
	// Selected subtitle language
	SubtitleLanguage string `json:"subtitleLanguage,omitempty"`
	// Current time in milliseconds
	Time *int64 `json:"time,omitempty"`
	// ID that is persisted across sessions to identify a browser
	UserId string `json:"userId,omitempty"`
	// Bitrate of the played back video rendition
	VideoBitrate *int64 `json:"videoBitrate,omitempty"`
	// Length of the video in milliseconds
	VideoDuration *int64 `json:"videoDuration,omitempty"`
	// ID of the video as configured via the analytics config
	VideoId string `json:"videoId,omitempty"`
	// Free form human readable video title as configured in the analytics config
	VideoTitle string `json:"videoTitle,omitempty"`
	// Resolution of the played back Video Rendition
	VideoPlaybackHeight *int32 `json:"videoPlaybackHeight,omitempty"`
	// Resolution of the played back Video Rendition
	VideoPlaybackWidth *int32 `json:"videoPlaybackWidth,omitempty"`
	// Time in milliseconds it took to start video playback
	VideoStartupTime *int64 `json:"videoStartupTime,omitempty"`
	// End time of the sample in the video (milliseconds)
	VideotimeEnd *int64 `json:"videotimeEnd,omitempty"`
	// Start time of the sample in the video (milliseconds)
	VideotimeStart *int64 `json:"videotimeStart,omitempty"`
	// Height of the video player on the page
	VideoWindowHeight *int32 `json:"videoWindowHeight,omitempty"`
	// Width of the video player on the page
	VideoWindowWidth *int32 `json:"videoWindowWidth,omitempty"`
}

