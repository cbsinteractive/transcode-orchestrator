package pagination

import(
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/serialization"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type MuxingsListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.Muxing `json:"items,omitempty"`
}

func (o *MuxingsListPagination) UnmarshalJSON(b []byte) error {
    var items []model.Muxing

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

    for _, i := range pageResp.Items {
        var base model.BaseMuxing
		serialization.Decode(i, &base)

        switch base.MuxingType() {
                case model.MuxingType_FMP4:
                    var v model.Fmp4Muxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_CMAF:
                    var v model.CmafMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_MP4:
                    var v model.Mp4Muxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_TS:
                    var v model.TsMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_WEBM:
                    var v model.WebmMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_MP3:
                    var v model.Mp3Muxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_PROGRESSIVE_WEBM:
                    var v model.ProgressiveWebmMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_PROGRESSIVE_MOV:
                    var v model.ProgressiveMovMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_PROGRESSIVE_TS:
                    var v model.ProgressiveTsMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_BROADCAST_TS:
                    var v model.BroadcastTsMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_CHUNKED_TEXT:
                    var v model.ChunkedTextMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_TEXT:
                    var v model.TextMuxing
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.MuxingType_SEGMENTED_RAW:
                    var v model.SegmentedRawMuxing
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

