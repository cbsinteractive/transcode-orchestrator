package hybrik

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
)

type packagingProtocol = string

const (
	packagingProtocolHLS  packagingProtocol = "hls"
	packagingProtocolDASH packagingProtocol = "dash"

	segmentationModeDefault           = "segmented_mp4"
	segmentLengthDefault              = uint(6)
	masterManifestFilenameDefaultTmpl = "master.%s"
)

type protocolCfg struct {
	extension, hybrikKind string
}

var supportedPackagingProtocols = map[packagingProtocol]protocolCfg{
	packagingProtocolHLS: {
		extension:  "m3u8",
		hybrikKind: "hls",
	},
	packagingProtocolDASH: {
		extension:  "mpd",
		hybrikKind: "dash",
	},
}

func (p *hybrikProvider) enrichCreateJobWithPackagingCfg(cj hybrik.CreateJob, jobCfg jobCfg,
	prevElements []hybrik.Element) (hybrik.CreateJob, error) {
	protocol := strings.ToLower(jobCfg.streamingParams.Protocol)

	protocolCfg, found := supportedPackagingProtocols[protocol]
	if !found {
		return cj, fmt.Errorf("protocol %q not supported", protocol)
	}

	segmentDuration := segmentLengthDefault
	if duration := jobCfg.streamingParams.SegmentDuration; duration > 0 {
		segmentDuration = duration
	}

	packagePayload := hybrik.PackagePayload{
		Location: p.transcodeLocationFrom(storageLocation{
			provider: jobCfg.destination.provider,
			path:     fmt.Sprintf("%s/%s", jobCfg.destination.path, protocol),
		}),
		FilePattern:        fmt.Sprintf(masterManifestFilenameDefaultTmpl, protocolCfg.extension),
		Kind:               protocolCfg.hybrikKind,
		SegmentationMode:   segmentationModeDefault,
		SegmentDurationSec: int(segmentDuration),
		ForceOriginalMedia: false,
	}

	switch protocol {
	case packagingProtocolHLS:
		packagePayload.HLS = &hybrik.HLSPackagingSettings{
			IncludeIFRAMEManifests: true,
			HEVCCodecIDPrefix:      h265VideoTagValueHVC1,
		}
	case packagingProtocolDASH:
		packagePayload.DASH = &hybrik.DASHPackagingSettings{
			UseSegmentList:     false,
			SegmentDurationSec: strconv.Itoa(int(segmentDuration)),
			SegmentationMode:   segmentationModeDefault,
		}
	}

	packagerUID := "packager"
	packageElement := hybrik.Element{
		UID:     packagerUID,
		Kind:    elementKindPackage,
		Payload: packagePayload,
	}

	if tag, found := jobCfg.computeTags[db.ComputeClassTranscodeDefault]; found {
		packageElement.Task = &hybrik.ElementTaskOptions{Tags: []string{tag}}
	}

	cj.Payload.Elements = append(cj.Payload.Elements, packageElement)

	manifestFromConnections := []hybrik.ConnectionFrom{}
	for _, task := range prevElements {
		manifestFromConnections = append(manifestFromConnections, hybrik.ConnectionFrom{Element: task.UID})
	}

	cj.Payload.Connections = append(cj.Payload.Connections,
		hybrik.Connection{
			From: manifestFromConnections,
			To:   hybrik.ConnectionTo{Success: []hybrik.ToSuccess{{Element: packagerUID}}},
		},
	)

	return cj, nil
}
