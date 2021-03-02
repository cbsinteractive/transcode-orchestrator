package hybrik

var RateControl = map[string]bool{
	"cbr": true,
	"vbr": true,
}

const (
	vbrVariabilityPercent = 10
)

func percentTargetOf(bitrate, percent int) int {
	return bitrate * (100 + percent) / 100
}
