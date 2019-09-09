package hybrik

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
)

type jobCfg struct {
	jobID                string
	destination          storageLocation
	sourceLocation       storageLocation
	source               hybrik.Element
	elementGroups        [][]hybrik.Element
	outputCfgs           map[string]outputCfg
	streamingParams      db.StreamingParams
	executionEnvironment db.ExecutionEnvironment
	executionFeatures    executionFeatures
	computeTags          map[db.ComputeClass]string
}

type outputCfg struct {
	localPreset db.Preset
	filename    string
}

const (
	assetContentsKindMetadata = "metadata"

	assetContentsStandardDolbyVisionMetadata = "dolbyvision_metadata"

	imfManifestExtension = ".xml"

	srcOptionResolveManifestKey = "resolve_manifest"
)

func (p *hybrikProvider) srcFrom(job *db.Job, src storageLocation) (hybrik.Element, error) {
	sourceAsset := p.assetPayloadFrom(src.provider, src.path, nil, job.ExecutionEnv)

	if strings.ToLower(filepath.Ext(src.path)) == imfManifestExtension {
		sourceAsset.Options = map[string]interface{}{
			srcOptionResolveManifestKey: true,
		}
	}

	assets := []hybrik.AssetPayload{sourceAsset}

	if sidecarLocation, ok := job.SidecarAssets[db.SidecarAssetKindDolbyVisionMetadata]; ok {
		sidecarStorageProvider, err := storageProviderFrom(sidecarLocation)
		if err != nil {
			return hybrik.Element{}, err
		}

		assets = append(assets, p.assetPayloadFrom(sidecarStorageProvider, sidecarLocation, []hybrik.AssetContents{{
			Kind: assetContentsKindMetadata,
			Payload: hybrik.AssetContentsPayload{
				Standard: assetContentsStandardDolbyVisionMetadata,
			},
		}}, job.ExecutionEnv))
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
	presets := map[string]outputCfg{}

	for _, output := range job.Outputs {
		presetName := output.Preset.Name
		presetResponse, err := p.GetPreset(presetName)
		if err != nil {
			return nil, err
		}

		localPreset, ok := presetResponse.(*db.LocalPreset)
		if !ok {
			return nil, fmt.Errorf("could not convert preset response into a db.LocalPreset")
		}

		presets[presetName] = outputCfg{
			localPreset: localPreset.Preset,
			filename:    output.FileName,
		}
	}

	return presets, nil
}

type elementAssembler func(jobCfg) ([][]hybrik.Element, error)

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
