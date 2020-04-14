package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type S3InputListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *S3InputListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
