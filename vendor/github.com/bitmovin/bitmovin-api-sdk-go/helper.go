package bitmovin

func StrPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int64Ptr(i int64) *int64 {
  return &i
}

func Float64Ptr(i float64) *float64 {
  return &i
}

func ToPtr(v interface{}) interface{} {
	return &v
}
