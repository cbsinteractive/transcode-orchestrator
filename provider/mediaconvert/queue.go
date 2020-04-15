package mediaconvert

import "github.com/cbsinteractive/video-transcoding-api/db"

func (p *mcProvider) canUsePreferredQueue(info db.SourceInfo) bool {
	return !p.requiresAcceleration(info)
}

const minSizeForAcceleration = 1_000_000_000

func (p *mcProvider) requiresAcceleration(info db.SourceInfo) bool {
	return info.FileSize > 0 && info.FileSize/minSizeForAcceleration >= 1
}
