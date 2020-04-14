package pagination

import(
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/serialization"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type OutputsListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.Output `json:"items,omitempty"`
}

func (o *OutputsListPagination) UnmarshalJSON(b []byte) error {
    var items []model.Output

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

    for _, i := range pageResp.Items {
        var base model.BaseOutput
		serialization.Decode(i, &base)

        switch base.OutputType() {
                case model.OutputType_AKAMAI_NETSTORAGE:
                    var v model.AkamaiNetStorageOutput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_AZURE:
                    var v model.AzureOutput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_GENERIC_S3:
                    var v model.GenericS3Output
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_GCS:
                    var v model.GcsOutput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_FTP:
                    var v model.FtpOutput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_LOCAL:
                    var v model.LocalOutput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_S3:
                    var v model.S3Output
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_S3_ROLE_BASED:
                    var v model.S3RoleBasedOutput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_SFTP:
                    var v model.SftpOutput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_AKAMAI_MSL:
                    var v model.AkamaiMslOutput
                    serialization.Decode(i, &v)
                    items = append(items, v)
                    break
                case model.OutputType_LIVE_MEDIA_INGEST:
                    var v model.LiveMediaIngestOutput
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

