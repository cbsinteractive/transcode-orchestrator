package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type Cea608CaptionInputStreamListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *Cea608CaptionInputStreamListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
