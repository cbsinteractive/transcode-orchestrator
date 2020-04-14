package pagination

import(
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/serialization"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type InputStreamsListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.InputStream `json:"items,omitempty"`
}

func (o *InputStreamsListPagination) UnmarshalJSON(b []byte) error {
    var items []model.InputStream

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

    for _, i := range pageResp.Items {
        var base model.BaseInputStream
		serialization.Decode(i, &base)

        switch base.InputStreamType() {
                case model.InputStreamType_INGEST:
                    var v model.IngestInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_CONCATENATION:
                    var v model.ConcatenationInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_TRIMMING_TIME_BASED:
                    var v model.TimeBasedTrimmingInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_TRIMMING_TIME_CODE_TRACK:
                    var v model.TimecodeTrackTrimmingInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_TRIMMING_H264_PICTURE_TIMING:
                    var v model.H264PictureTimingTrimmingInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_AUDIO_MIX:
                    var v model.AudioMixInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_SIDECAR_DOLBY_VISION_METADATA:
                    var v model.DolbyVisionMetadataIngestInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_CAPTION_CEA608:
                    var v model.Cea608CaptionInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_CAPTION_CEA708:
                    var v model.Cea708CaptionInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_FILE:
                    var v model.FileInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputStreamType_DVB_TELETEXT:
                    var v model.DvbTeletextInputStream
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                default:
                    items = append(items, base)
        }
    }

    o.TotalCount = pageResp.TotalCount
    o.Offset = pageResp.Offset
    o.Limit = pageResp.Limit
    o.Previous = pageResp.Previous
    o.Next = pageResp.Next
    o.Items = items
    return nil
}

