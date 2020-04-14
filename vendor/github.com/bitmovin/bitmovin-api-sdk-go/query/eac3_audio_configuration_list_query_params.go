package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type Eac3AudioConfigurationListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *Eac3AudioConfigurationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
