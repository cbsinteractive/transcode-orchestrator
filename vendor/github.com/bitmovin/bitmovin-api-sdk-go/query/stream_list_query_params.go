package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type StreamListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *StreamListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
