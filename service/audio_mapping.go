package service

import (
	"fmt"

	"github.com/cbsinteractive/transcode-orchestrator/job"
)

type downmixMap map[job.ChannelLayout]map[int]bool

var downmixMappings = map[int]downmixMap{
	2: {
		job.ChannelLayoutLFE:           nil,
		job.ChannelLayoutCenter:        {0: true, 1: true},
		job.ChannelLayoutLeft:          {0: true},
		job.ChannelLayoutRight:         {1: true},
		job.ChannelLayoutLeftSurround:  {0: true},
		job.ChannelLayoutRightSurround: {1: true},
		job.ChannelLayoutLeftBack:      {0: true},
		job.ChannelLayoutRightBack:     {1: true},
		job.ChannelLayoutLeftTotal:     {0: true},
		job.ChannelLayoutRightTotal:    {1: true},
	},
}

//AudioDownmixMapping converts
func AudioDownmixMapping(ad job.Downmix) ([][]bool, error) {
	dm, found := downmixMappings[len(ad.Dst)]
	if !found {
		return nil, fmt.Errorf("no downmix config found when converting %d src channels to %d destination channels",
			len(ad.Src), len(ad.Dst))
	}

	m := make([][]bool, len(ad.Dst))
	for i := range m {
		m[i] = make([]bool, len(ad.Src))
	}

	for srcIdx, src := range ad.Src {
		for destIdx, enabled := range dm[job.ChannelLayout(src.Layout)] {
			m[destIdx][srcIdx] = enabled
		}
	}

	return m, nil
}
