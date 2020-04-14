package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type S3RoleBasedOutputListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *S3RoleBasedOutputListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
