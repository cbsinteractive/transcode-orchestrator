package pagination

import(
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/serialization"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type FiltersListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.Filter `json:"items,omitempty"`
}

func (o *FiltersListPagination) UnmarshalJSON(b []byte) error {
    var items []model.Filter

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

    for _, i := range pageResp.Items {
        var base model.BaseFilter
		serialization.Decode(i, &base)

        switch base.FilterType() {
                case model.FilterType_CROP:
                    var v model.CropFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_CONFORM:
                    var v model.ConformFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_WATERMARK:
                    var v model.WatermarkFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_ENHANCED_WATERMARK:
                    var v model.EnhancedWatermarkFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_ROTATE:
                    var v model.RotateFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_DEINTERLACE:
                    var v model.DeinterlaceFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_AUDIO_MIX:
                    var v model.AudioMixFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_DENOISE_HQDN3_D:
                    var v model.DenoiseHqdn3dFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_TEXT:
                    var v model.TextFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_UNSHARP:
                    var v model.UnsharpFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_SCALE:
                    var v model.ScaleFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_INTERLACE:
                    var v model.InterlaceFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_AUDIO_VOLUME:
                    var v model.AudioVolumeFilter
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.FilterType_EBU_R128_SINGLE_PASS:
                    var v model.EbuR128SinglePassFilter
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

