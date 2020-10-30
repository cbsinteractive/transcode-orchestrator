package hybrik

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/db"
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

type outputCfg = db.TranscodeOutput

const (
	assetContentsKindMetadata = "metadata"

	assetContentsStandardDolbyVisionMetadata = "dolbyvision_metadata"

	imfManifestExtension = ".xml"

	srcOptionResolveManifestKey = "resolve_manifest"
)

func (p *hybrikProvider) srcFrom(job *db.Job, src storageLocation) (hybrik.Element, error) {
	sourceAsset := p.assetPayloadFrom(src.provider, src.path, nil, job.ExecutionEnv.InputAlias)

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
		}}, job.ExecutionEnv.InputAlias))
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

func (p *hybrikProvider) outputCfgsFrom(ctx context.Context, job *db.Job) (map[string]outputCfg, error) {
	presets := map[string]outputCfg{}

	for _, output := range job.Outputs {
		presets[output.Preset.Name] = output
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
