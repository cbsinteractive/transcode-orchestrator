package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type SpekeDrmListQueryParams struct {
    Offset string `query:"offset"`
    Limit string `query:"limit"`
}

func (q *SpekeDrmListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
