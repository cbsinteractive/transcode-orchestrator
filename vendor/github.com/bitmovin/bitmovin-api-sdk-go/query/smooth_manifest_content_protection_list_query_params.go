package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type SmoothManifestContentProtectionListQueryParams struct {
    Offset int64 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *SmoothManifestContentProtectionListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
