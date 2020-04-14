package model

type StreamInfosDetails struct {
	// The id of the stream (required)
	Id *int64 `json:"id,omitempty"`
	// The media type of the stream (required)
	MediaType MediaType `json:"mediaType,omitempty"`
	// The width of the stream, if it is a video stream
	Width *int32 `json:"width,omitempty"`
	// The height of the stream, if it is a video stream
	Height *int32 `json:"height,omitempty"`
	// The rate (sample rate / fps) of the stream (required)
	Rate *int64 `json:"rate,omitempty"`
	// The codec of the input stream (required)
	Codec LiveEncodingCodec `json:"codec,omitempty"`
	// The minimum samples read per second within the last minute (required)
	SamplesReadPerSecondMin *float64 `json:"samplesReadPerSecondMin,omitempty"`
	// The maximum samples read per second within the last minute (required)
	SamplesReadPerSecondMax *float64 `json:"samplesReadPerSecondMax,omitempty"`
	// The average samples read per second within the last minute (required)
	SamplesReadPerSecondAvg *float64 `json:"samplesReadPerSecondAvg,omitempty"`
	// The minimum amount of backup samples used per second within the last minute. This will be written when no live stream is ingested. The last picture will be repeated with silent audio. (required)
	SamplesBackupPerSecondMin *float64 `json:"samplesBackupPerSecondMin,omitempty"`
	// The maximum amount of backup samples used per second within the last minute. This will be written when no live stream is ingested. The last picture will be repeated with silent audio. (required)
	SamplesBackupPerSecondMax *float64 `json:"samplesBackupPerSecondMax,omitempty"`
	// The average amount of backup samples used per second within the last minute. This will be written when no live stream is ingested. The last picture will be repeated with silent audio. (required)
	SamplesBackupPerSecondAvg *float64 `json:"samplesBackupPerSecondAvg,omitempty"`
	// The minimum bytes read per second within the last minute (required)
	BytesReadPerSecondMin *float64 `json:"bytesReadPerSecondMin,omitempty"`
	// The maximum bytes read per second within the last minute (required)
	BytesReadPerSecondMax *float64 `json:"bytesReadPerSecondMax,omitempty"`
	// The average bytes read per second within the last minute (required)
	BytesReadPerSecondAvg *float64 `json:"bytesReadPerSecondAvg,omitempty"`
	// The minimum amount of backup bytes used per second within the last minute. This will be written when no live stream is ingested. The last picture will be repeated with silent audio. (required)
	BytesBackupPerSecondMin *float64 `json:"bytesBackupPerSecondMin,omitempty"`
	// The maximum amount of backup bytes used per second within the last minute. This will be written when no live stream is ingested. The last picture will be repeated with silent audio. (required)
	BytesBackupPerSecondMax *float64 `json:"bytesBackupPerSecondMax,omitempty"`
	// The average amount of backup bytes used per second within the last minute. This will be written when no live stream is ingested. The last picture will be repeated with silent audio. (required)
	BytesBackupPerSecondAvg *float64 `json:"bytesBackupPerSecondAvg,omitempty"`
}

