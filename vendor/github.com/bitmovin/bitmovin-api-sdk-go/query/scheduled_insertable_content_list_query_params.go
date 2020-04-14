package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type ScheduledInsertableContentListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *ScheduledInsertableContentListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
