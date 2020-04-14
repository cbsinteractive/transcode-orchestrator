package model

type AnalyticsMaxConcurrentViewersResponse struct {
	Rows []float64 `json:"rows,omitempty"`
	// Number of rows returned
	RowCount *int64 `json:"rowCount,omitempty"`
	ColumnLabels []AnalyticsColumnLabel `json:"columnLabels,omitempty"`
}

