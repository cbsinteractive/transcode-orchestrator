package legacy

import (
	"time"

	"github.com/cbsinteractive/pkg/timecode"
	"github.com/cbsinteractive/transcode-orchestrator/av"
)

func MigrateV1toV2(o Job, op ...Preset) (n av.Job) {
	opmap := map[string]Preset{}
	for _, op := range op {
		opmap[op.Name] = op
	}

	n.ID = o.JobID
	n.Name = o.Name
	n.CreatedAt = o.CreationTime

	n.Env.Cloud = o.ExecutionEnv.Cloud
	n.Env.Region = o.ExecutionEnv.Region
	n.Env.InputAlias = o.ExecutionEnv.InputAlias
	n.Env.OutputAlias = o.ExecutionEnv.OutputAlias
	n.Env.Tags = o.ExecutionEnv.ComputeTags

	n.Features = o.ExecutionFeatures
	n.ExtraFiles = o.SidecarAssets
	n.Labels = o.Labels

	n.Input.Name = o.Source
	n.Input.Splice = o.SourceSplice
	n.Input.Size = o.SourceInfo.FileSize
	n.Input.Video.Width = o.SourceInfo.Width
	n.Input.Video.Height = o.SourceInfo.Height
	//	n.Input.Video.Framerate = o.SourceInfo.FrameRate
	n.Input.Video.Scantype = o.SourceInfo.ScanType
	n.Input.ExplicitKeyframeOffsets = o.ExplicitKeyframeOffsets

	nmix := av.Downmix{}
	for _, oc := range o.AudioDownmix.SrcChannels {
		nmix.Src = append(nmix.Src, av.AudioChannel{
			TrackIdx:   oc.TrackIdx,
			ChannelIdx: oc.ChannelIdx,
			Layout:     oc.Layout,
		})
	}
	for _, oc := range o.AudioDownmix.DestChannels {
		nmix.Dst = append(nmix.Dst, av.AudioChannel{
			TrackIdx:   oc.TrackIdx,
			ChannelIdx: oc.ChannelIdx,
			Layout:     oc.Layout,
		})
	}

	n.Output.Path = o.DestinationBasePath
	for _, of := range o.Outputs {
		p := opmap[of.PresetMap.Name]
		n.Output.Add(av.File{
			Name: of.FileName,
			//	Container: op.Container,
			Video: av.Video{
				Codec:    p.Video.Codec,
				Profile:  p.Video.Profile,
				Level:    p.Video.ProfileLevel,
				Width:    p.Video.Width,
				Height:   p.Video.Height,
				Scantype: p.Video.InterlaceMode,
			},
			Downmix: &nmix,
		})
	}

	return n
}

type Job struct {
	JobID, Name, ProviderName, ProviderJobID string
	CreationTime                             time.Time
	ExecutionEnv                             struct {
		Cloud, Region, InputAlias, OutputAlias string
		ComputeTags                            map[string]string
	}
	ExecutionFeatures map[string]interface{}
	SidecarAssets     map[string]string
	Labels            []string

	Source       string
	SourceSplice timecode.Splice
	SourceInfo   struct {
		Width     int
		Height    int
		FrameRate float64
		FileSize  int64
		ScanType  string
	}
	DestinationBasePath string
	Outputs             []struct {
		FileName  string
		PresetMap struct {
			Name            string
			ProviderMapping map[string]string
			Output          struct {
				Extension string
			}
		}
	}
	AudioDownmix struct {
		SrcChannels, DestChannels []struct {
			TrackIdx, ChannelIdx int
			Layout               string
		}
	}
	ExplicitKeyframeOffsets []float64
}

type Preset struct {
	Name, Description, SourceContainer, Container, RateControl string
	TwoPass                                                    bool
	Video                                                      struct {
		Codec, Profile, ProfileLevel string
		Width                        int     `json:",string"`
		Height                       int     `json:",string"`
		Bitrate                      int     `json:",string"`
		GopSize                      float64 `json:",string"`
		GopUnit                      string
		GopMode                      string
		InterlaceMode                string
		HDR10                        struct {
			Enabled         bool
			MaxCLL, MaxFALL int
			MasterDisplay   string
		}
		DolbyVisionSettings struct {
			Enabled bool
		}
		Overlays struct {
			Images         []struct{ URL string }
			TimecodeBurnin struct {
				Enabled  bool
				FontSize int
				Position int
				Prefix   string
			}
		}

		Crop struct {
			Left, Top, Right, Bottom int
		}
		Framerate struct {
			Numerator   int
			Denominator int
		}
		Audio struct {
			Codec          string
			Bitrate        int `json:",string"`
			Normalization  bool
			DiscreteTracks bool
		}
	}
}
