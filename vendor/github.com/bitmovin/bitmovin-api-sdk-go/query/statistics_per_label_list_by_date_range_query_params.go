package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type StatisticsPerLabelListByDateRangeQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Labels string `query:"labels"`
}

func (q *StatisticsPerLabelListByDateRangeQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
