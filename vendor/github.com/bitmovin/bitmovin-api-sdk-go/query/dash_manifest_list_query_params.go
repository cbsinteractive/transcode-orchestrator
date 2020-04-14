package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type DashManifestListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    EncodingId string `query:"encodingId"`
}

func (q *DashManifestListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
