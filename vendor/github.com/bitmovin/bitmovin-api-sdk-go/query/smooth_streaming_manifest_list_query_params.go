package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type SmoothStreamingManifestListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    EncodingId string `query:"encodingId"`
}

func (q *SmoothStreamingManifestListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
