package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type Cea708CaptionInputStreamListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *Cea708CaptionInputStreamListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
