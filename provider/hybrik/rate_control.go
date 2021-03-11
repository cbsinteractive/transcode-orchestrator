package hybrik

import "github.com/cbsinteractive/transcode-orchestrator/client/transcoding/job"

const (
	vbrVariability = 10
)

var RateControl = map[string]int{
	"vbr": 1,
	"cbr": 0,
}

func percentTarget(b job.Bitrate, percent int) int {
	on := RateControl[canon(b.Control)]
	return on * b.BPS * (100 + percent) / 100
}
