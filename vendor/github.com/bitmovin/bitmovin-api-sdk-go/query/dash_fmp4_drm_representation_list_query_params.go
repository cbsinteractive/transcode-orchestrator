package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type DashFmp4DrmRepresentationListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *DashFmp4DrmRepresentationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
