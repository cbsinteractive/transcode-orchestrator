package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type DailyStatisticsListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *DailyStatisticsListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
