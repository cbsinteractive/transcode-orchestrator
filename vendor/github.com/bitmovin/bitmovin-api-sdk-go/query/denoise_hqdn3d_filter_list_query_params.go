package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type DenoiseHqdn3dFilterListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *DenoiseHqdn3dFilterListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
