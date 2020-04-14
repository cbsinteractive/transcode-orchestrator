package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type EncodingListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Sort string `query:"sort"`
    Type_ string `query:"type"`
    Status string `query:"status"`
    CloudRegion model.CloudRegion `query:"cloudRegion"`
    EncoderVersion string `query:"encoderVersion"`
    Name string `query:"name"`
    Search string `query:"search"`
}

func (q *EncodingListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
