package model
import (
	"time"
)

type StreamInfo struct {
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
	// It MUST match the value of the GROUP-ID attribute of an Audio EXT-X-MEDIA tag elsewhere in the Master Playlist. Either this or `audioGroups` must be set.
	Audio string `json:"audio,omitempty"`
	// HLS Audio Group Configuration. You will want to use this configuration property in case you specify conditions on audio streams. The first matching audio group will be used for the specific variant stream. Either this or `audio` must be set.
	AudioGroups *AudioGroupConfiguration `json:"audioGroups,omitempty"`
	// It MUST match the value of the GROUP-ID attribute of a Video EXT-X-MEDIA tag elsewhere in the Master Playlist
	Video string `json:"video,omitempty"`
	// It MUST match the value of the GROUP-ID attribute of a Subtitles EXT-X-MEDIA tag elsewhere in the Master Playlist
	Subtitles string `json:"subtitles,omitempty"`
	// If the value is not 'NONE', it MUST match the value of the GROUP-ID attribute of a Closed Captions EXT-X-MEDIA tag elsewhere in the Playlist (required)
	ClosedCaptions string `json:"closedCaptions,omitempty"`
	// Id of the encoding. (required)
	EncodingId string `json:"encodingId,omitempty"`
	// Id of the stream. (required)
	StreamId string `json:"streamId,omitempty"`
	// Id of the muxing. (required)
	MuxingId string `json:"muxingId,omitempty"`
	// Id of the DRM.
	DrmId string `json:"drmId,omitempty"`
	// Path to segments. (required)
	SegmentPath string `json:"segmentPath,omitempty"`
	// The URI of the playlist file. (required)
	Uri string `json:"uri,omitempty"`
	// Number of the first segment. Default is 0.
	StartSegmentNumber *int64 `json:"startSegmentNumber,omitempty"`
	// Number of the last segment. Default is the last one that was encoded.
	EndSegmentNumber *int64 `json:"endSegmentNumber,omitempty"`
	// Force the addition of the frame rate attribute to all stream infos.
	ForceFrameRateAttribute *bool `json:"forceFrameRateAttribute,omitempty"`
	// Force the addition of the video-range attribute to all stream infos.
	ForceVideoRangeAttribute *bool `json:"forceVideoRangeAttribute,omitempty"`
}

