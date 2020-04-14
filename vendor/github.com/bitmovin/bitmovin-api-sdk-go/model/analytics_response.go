package model

type AnalyticsResponse struct {
	Rows []map[string]interface{} `json:"rows,omitempty"`
	// Number of rows returned
	RowCount *int64 `json:"rowCount,omitempty"`
	ColumnLabels []AnalyticsColumnLabel `json:"columnLabels,omitempty"`
}

