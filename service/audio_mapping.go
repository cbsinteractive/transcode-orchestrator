package service

import (
	"fmt"

	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type downmixMap map[db.ChannelLayout]map[int]bool

var downmixMappings = map[int]downmixMap{
	2: {
		db.ChannelLayoutLFE:           nil,
		db.ChannelLayoutCenter:        {0: true, 1: true},
		db.ChannelLayoutLeft:          {0: true},
		db.ChannelLayoutRight:         {1: true},
		db.ChannelLayoutLeftSurround:  {0: true},
		db.ChannelLayoutRightSurround: {1: true},
		db.ChannelLayoutLeftBack:      {0: true},
		db.ChannelLayoutRightBack:     {1: true},
		db.ChannelLayoutLeftTotal:     {0: true},
		db.ChannelLayoutRightTotal:    {1: true},
	},
}

//AudioDownmixMapping converts
func AudioDownmixMapping(ad db.AudioDownmix) ([][]bool, error) {
	dm, found := downmixMappings[len(ad.DestChannels)]
	if !found {
		return nil, fmt.Errorf("no downmix config found when converting %d src channels to %d destination channels",
			len(ad.SrcChannels), len(ad.DestChannels))
	}

	m := make([][]bool, len(ad.DestChannels))
	for i := range m {
		m[i] = make([]bool, len(ad.SrcChannels))
	}

	for srcIdx, src := range ad.SrcChannels {
		for destIdx, enabled := range dm[db.ChannelLayout(src.Layout)] {
			m[destIdx][srcIdx] = enabled
		}
	}

	return m, nil
}
