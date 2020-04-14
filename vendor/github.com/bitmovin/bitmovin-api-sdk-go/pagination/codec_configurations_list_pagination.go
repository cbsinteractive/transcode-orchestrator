package pagination

import(
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/serialization"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type CodecConfigurationsListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.CodecConfiguration `json:"items,omitempty"`
}

func (o *CodecConfigurationsListPagination) UnmarshalJSON(b []byte) error {
    var items []model.CodecConfiguration

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

    for _, i := range pageResp.Items {
        var base model.BaseCodecConfiguration
		serialization.Decode(i, &base)

        switch base.CodecConfigType() {
                case model.CodecConfigType_AAC:
                    var v model.AacAudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_HE_AAC_V1:
                    var v model.HeAacV1AudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_HE_AAC_V2:
                    var v model.HeAacV2AudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_H264:
                    var v model.H264VideoConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_H265:
                    var v model.H265VideoConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_VP9:
                    var v model.Vp9VideoConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_VP8:
                    var v model.Vp8VideoConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_MP2:
                    var v model.Mp2AudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_MP3:
                    var v model.Mp3AudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_AC3:
                    var v model.Ac3AudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_EAC3:
                    var v model.Eac3AudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_OPUS:
                    var v model.OpusAudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_VORBIS:
                    var v model.VorbisAudioConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_MJPEG:
                    var v model.MjpegVideoConfiguration
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.CodecConfigType_AV1:
                    var v model.Av1VideoConfiguration
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

