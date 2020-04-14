package pagination

import(
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/serialization"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type DrmsListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.Drm `json:"items,omitempty"`
}

func (o *DrmsListPagination) UnmarshalJSON(b []byte) error {
    var items []model.Drm

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

    for _, i := range pageResp.Items {
        var base model.BaseDrm
		serialization.Decode(i, &base)

        switch base.DrmType() {
                case model.DrmType_WIDEVINE:
                    var v model.WidevineDrm
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.DrmType_PLAYREADY:
                    var v model.PlayReadyDrm
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.DrmType_PRIMETIME:
                    var v model.PrimeTimeDrm
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.DrmType_FAIRPLAY:
                    var v model.FairPlayDrm
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.DrmType_MARLIN:
                    var v model.MarlinDrm
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.DrmType_CLEARKEY:
                    var v model.ClearKeyDrm
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.DrmType_AES:
                    var v model.AesEncryptionDrm
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.DrmType_CENC:
                    var v model.CencDrm
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.DrmType_SPEKE:
                    var v model.SpekeDrm
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

