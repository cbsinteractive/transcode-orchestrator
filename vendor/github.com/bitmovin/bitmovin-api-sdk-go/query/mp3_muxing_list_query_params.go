package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type Mp3MuxingListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *Mp3MuxingListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
