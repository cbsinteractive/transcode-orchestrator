package job

import (
	"fmt"
)

type downmixMap map[ChannelLayout]map[int]bool

var downmixMappings = map[int]downmixMap{
	2: {
		ChannelLayoutLFE:           nil,
		ChannelLayoutCenter:        {0: true, 1: true},
		ChannelLayoutLeft:          {0: true},
		ChannelLayoutRight:         {1: true},
		ChannelLayoutLeftSurround:  {0: true},
		ChannelLayoutRightSurround: {1: true},
		ChannelLayoutLeftBack:      {0: true},
		ChannelLayoutRightBack:     {1: true},
		ChannelLayoutLeftTotal:     {0: true},
		ChannelLayoutRightTotal:    {1: true},
	},
}

//AudioDownmixMapping converts
func AudioDownmixMapping(ad *Downmix) ([][]bool, error) {
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
		for destIdx, enabled := range dm[ChannelLayout(src.Layout)] {
			m[destIdx][srcIdx] = enabled
		}
	}

	return m, nil
}
