package mediaconvert

import "github.com/cbsinteractive/transcode-orchestrator/db"

func (p *driver) canUsePreferredQueue(info db.File) bool {
	return !p.requiresAcceleration(info)
}

const minSizeForAcceleration = 1_000_000_000

func (p *driver) requiresAcceleration(info db.File) bool {
	return false // hack: (ts) temporarily disabled this due to bugs in EMC (9/JUNE/2020)
	//return info.FileSize > 0 && info.FileSize/minSizeForAcceleration >= 1
}
