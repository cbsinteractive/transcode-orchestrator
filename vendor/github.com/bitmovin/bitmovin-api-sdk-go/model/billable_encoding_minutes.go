package model

type BillableEncodingMinutes struct {
	EncodingMode EncodingMode `json:"encodingMode,omitempty"`
	Codec CodecConfigType `json:"codec,omitempty"`
	PerTitleResultStream StatisticsPerTitleStream `json:"perTitleResultStream,omitempty"`
	PsnrMode PsnrPerStreamMode `json:"psnrMode,omitempty"`
	BillableMinutes *BillableEncodingMinutesDetails `json:"billableMinutes,omitempty"`
}

