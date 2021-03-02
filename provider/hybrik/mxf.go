package hybrik

import (
	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func applyMXF(t *hy.TranscodeTarget, f job.File) bool {
	switch {
	case f.Video.DolbyVision.Enabled:
		t.NumPasses = 2
	case f.Video.HDR10.Enabled:
		// overriding this value to tell hybrik where to fetch the HDR metadata
		t.Video.HDR10.Source = "source_metadata"
		t.NumPasses = 2
	default:
		return false
	}
	return true
}
