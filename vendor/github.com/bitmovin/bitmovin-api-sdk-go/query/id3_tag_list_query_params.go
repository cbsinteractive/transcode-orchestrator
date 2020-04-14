package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type Id3TagListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *Id3TagListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
