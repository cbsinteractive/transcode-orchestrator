package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type H264PictureTimingTrimmingInputStreamListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *H264PictureTimingTrimmingInputStreamListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
