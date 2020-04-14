package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type PlaintextId3TagListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *PlaintextId3TagListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
