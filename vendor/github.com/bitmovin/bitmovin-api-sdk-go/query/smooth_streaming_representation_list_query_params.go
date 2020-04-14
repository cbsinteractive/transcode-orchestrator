package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type SmoothStreamingRepresentationListQueryParams struct {
    Offset int64 `query:"offset"`
    Limit int64 `query:"limit"`
}

func (q *SmoothStreamingRepresentationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
