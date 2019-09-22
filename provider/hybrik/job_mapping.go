package hybrik

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/pkg/errors"
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
	transcodeCfg db.TranscodeOutputConfig
	filename     string
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

	for i, output := range job.Outputs {
		if len(output.Preset.ProviderMapping) > 0 {
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
				transcodeCfg: localPreset.Preset,
				filename:     output.FileName,
			}
		} else if cfg := output.Config; cfg != nil {
			presets[fmt.Sprintf("output_%d", i)] = outputCfg{
				transcodeCfg: *cfg,
				filename:     output.FileName,
			}
		} else {
			return nil, errors.New("all outputs must either reference a stored preset or include an output config")
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
		_ = p.dolbyVisionElementAssembler // switch back to this once Hybrik fixes bug with GCP jobs hanging
		return p.dolbyVisionLegacyElementAssembler, nil
	}

	return p.defaultElementAssembler, nil
}
