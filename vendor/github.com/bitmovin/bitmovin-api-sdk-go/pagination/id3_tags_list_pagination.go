package pagination

import(
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/serialization"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type Id3TagsListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.Id3Tag `json:"items,omitempty"`
}

func (o *Id3TagsListPagination) UnmarshalJSON(b []byte) error {
    var items []model.Id3Tag

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

    for _, i := range pageResp.Items {
        var base model.BaseId3Tag
		serialization.Decode(i, &base)

        switch base.Id3TagType() {
                case model.Id3TagType_RAW:
                    var v model.RawId3Tag
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.Id3TagType_FRAME_ID:
                    var v model.FrameIdId3Tag
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.Id3TagType_PLAIN_TEXT:
                    var v model.PlaintextId3Tag
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

