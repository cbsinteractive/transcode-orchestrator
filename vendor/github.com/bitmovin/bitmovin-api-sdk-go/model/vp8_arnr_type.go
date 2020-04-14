package model
// Vp8ArnrType : altref noise reduction filter type
type Vp8ArnrType string

// List of Vp8ArnrType
const (
	Vp8ArnrType_BACKWARD Vp8ArnrType = "BACKWARD"
	Vp8ArnrType_FORWARD Vp8ArnrType = "FORWARD"
	Vp8ArnrType_CENTERED Vp8ArnrType = "CENTERED"
)