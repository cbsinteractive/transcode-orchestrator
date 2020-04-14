package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type NotificationStateEntryListByNotificationIdQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *NotificationStateEntryListByNotificationIdQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
