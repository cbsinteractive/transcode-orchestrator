package hybrik

import (
	"fmt"

	"github.com/NYTimes/video-transcoding-api/db"
	"github.com/NYTimes/video-transcoding-api/provider"
	"github.com/cbsinteractive/hybrik-sdk-go"
)

type jobCfg struct {
	jobID             string
	destination       storageLocation
	sourceLocation    storageLocation
	source            hybrik.Element
	tasks             []hybrik.Element
	outputCfgs        map[string]outputCfg
	streamingParams   db.StreamingParams
	executionFeatures executionFeatures
	computeTags       map[db.ComputeClass]string
}

type outputCfg struct {
	preset   hybrik.Preset
	filename string
}

const (
	assetContentsKindMetadata = "metadata"

	assetContentsStandardDolbyVisionMetadata = "dolbyvision_metadata"
)

func srcFrom(job *db.Job, src storageLocation) (hybrik.Element, error) {
	assets := []hybrik.AssetPayload{
		{
			StorageProvider: src.provider,
			URL:             src.path,
		},
	}

	if sidecarLocation, ok := job.SidecarAssets[db.SidecarAssetKindDolbyVisionMetadata]; ok {
		sidecarStorageProvider, err := storageProviderFrom(sidecarLocation)
		if err != nil {
			return hybrik.Element{}, err
		}

		assets = append(assets, hybrik.AssetPayload{
			StorageProvider: sidecarStorageProvider,
			URL:             sidecarLocation,
			Contents: []hybrik.AssetContents{{
				Kind: assetContentsKindMetadata,
				Payload: hybrik.AssetContentsPayload{
					Standard: assetContentsStandardDolbyVisionMetadata,
				},
			}},
		})
	}

	return hybrik.Element{
		UID:  "source_file",
		Kind: elementKindSource,
		Payload: hybrik.ElementPayload{
			Kind:    "asset_urls",
			Payload: assets,
		},
	}, nil
}

func (p *hybrikProvider) outputCfgsFrom(job *db.Job) (map[string]outputCfg, error) {
	presetCh := make(chan *presetResult)
	presets := map[string]outputCfg{}

	for _, output := range job.Outputs {
		presetID, ok := output.Preset.ProviderMapping[Name]
		if !ok {
			return nil, provider.ErrPresetMapNotFound
		}

		if _, ok := presets[presetID]; ok {
			continue
		}

		presets[presetID] = outputCfg{filename: output.FileName}

		go p.makeGetPresetRequest(presetID, presetCh)
	}

	for i := 0; i < len(presets); i++ {
		res := <-presetCh
		err, isErr := res.preset.(error)
		if isErr {
			return nil, fmt.Errorf("error getting preset info: %s", err)
		}

		preset, ok := res.preset.(hybrik.Preset)
		if !ok {
			return nil, fmt.Errorf("preset content was not a hybrik.Preset: %+v", res.preset)
		}

		stub, ok := presets[res.presetID]
		if !ok {
			return nil, fmt.Errorf("could not find stubbed outputCfg for preset with id %q", res.presetID)
		}

		stub.preset = preset
		presets[res.presetID] = stub
	}

	return presets, nil
}

type presetResult struct {
	presetID string
	preset   interface{}
}

func (p *hybrikProvider) makeGetPresetRequest(presetID string, ch chan *presetResult) {
	presetOutput, err := p.GetPreset(presetID)
	result := new(presetResult)
	result.presetID = presetID
	if err != nil {
		result.preset = err
		ch <- result
	} else {
		result.preset = presetOutput
		ch <- result
	}
}

type elementAssembler func(jobCfg) ([]hybrik.Element, error)

func (p *hybrikProvider) elementAssemblerFrom(cfgs map[string]outputCfg) (elementAssembler, error) {
	dolbyVisionEnabled, err := dolbyVisionEnabledOnAllPresets(cfgs)
	if err != nil {
		return nil, err
	}

	if dolbyVisionEnabled {
		return p.dolbyVisionElementAssembler, nil
	}

	return p.defaultElementAssembler, nil
}
