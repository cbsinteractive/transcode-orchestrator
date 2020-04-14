package pagination

import(
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/serialization"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type InputsListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.Input `json:"items,omitempty"`
}

func (o *InputsListPagination) UnmarshalJSON(b []byte) error {
    var items []model.Input

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

    for _, i := range pageResp.Items {
        var base model.BaseInput
		serialization.Decode(i, &base)

        switch base.InputType() {
                case model.InputType_AKAMAI_NETSTORAGE:
                    var v model.AkamaiNetStorageInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_ASPERA:
                    var v model.AsperaInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_AZURE:
                    var v model.AzureInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_REDUNDANT_RTMP:
                    var v model.RedundantRtmpInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_FTP:
                    var v model.FtpInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_GENERIC_S3:
                    var v model.GenericS3Input
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_GCS:
                    var v model.GcsInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_HTTP:
                    var v model.HttpInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_HTTPS:
                    var v model.HttpsInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_LOCAL:
                    var v model.LocalInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_RTMP:
                    var v model.RtmpInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_S3:
                    var v model.S3Input
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_S3_ROLE_BASED:
                    var v model.S3RoleBasedInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_SFTP:
                    var v model.SftpInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_TCP:
                    var v model.TcpInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_UDP:
                    var v model.UdpInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_UDP_MULTICAST:
                    var v model.UdpMulticastInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_ZIXI:
                    var v model.ZixiInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.InputType_SRT:
                    var v model.SrtInput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                default:
                    items = append(items, base)
        }
    }

    o.TotalCount = pageResp.TotalCount
    o.Offset = pageResp.Offset
    o.Limit = pageResp.Limit
    o.Previous = pageResp.Previous
    o.Next = pageResp.Next
    o.Items = items
    return nil
}

