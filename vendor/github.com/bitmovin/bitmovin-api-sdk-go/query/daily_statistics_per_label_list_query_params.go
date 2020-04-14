package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type DailyStatisticsPerLabelListQueryParams struct {
    Offset int64 `query:"offset"`
    Limit int64 `query:"limit"`
    Labels string `query:"labels"`
}

func (q *DailyStatisticsPerLabelListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
