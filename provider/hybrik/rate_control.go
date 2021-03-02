package hybrik

const (
	vbrVariabilityPercent = 10
)

var RateControl = map[string]bool{
	"cbr": true,
	"vbr": true,
}

func percentTargetOf(bitrate, percent int) int {
	return bitrate * (100 + percent) / 100
}
