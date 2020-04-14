package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type EbuR128SinglePassFilterListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *EbuR128SinglePassFilterListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
