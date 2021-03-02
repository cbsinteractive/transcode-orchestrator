package hybrik

import (
	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

const (
	assetContentsKindMetadata = "metadata"
	imfManifestExtension      = ".xml"

	srcOptionResolveManifestKey = "resolve_manifest"
)

type jobCfg struct {
	jobID string
	//destination          storageLocation
	//sourceLocation       storageLocation
	out job.Dir

	source               hybrik.Element
	elementGroups        [][]hybrik.Element
	streaming            job.Streaming
	executionEnvironment job.ExecutionEnvironment
	executionFeatures    executionFeatures
	computeTags          map[job.ComputeClass]string
}

const SourceUID = "source_file"

func (p *hybrikProvider) srcFrom(j *Job) (hybrik.Element, error) {
	creds := j.ExecutionEnv.InputAlias
	assets := []hybrik.AssetPayload{p.asset(j.Input, creds)}

	if dolby := j.Asset(job.SidecarAssetKindDolbyVisionMetadata); dolby != nil {
		assets = append(assets, p.asset(dolby, []hybrik.AssetContents{{
			Kind:    "metadata",
			Payload: hybrik.AssetContentsPayload{Standard: "dolbyvision_metadata"},
		}}, creds))
	}
	return hybrik.Element{
		UID:  SourceUID,
		Kind: elementKindSource,
		Payload: hybrik.ElementPayload{
			Kind:    "asset_urls",
			Payload: assets,
		},
	}
}

type elementAssembler func(jobCfg) ([][]hybrik.Element, error)

func (p *hybrikProvider) elementAssemblerFrom(d *job.Dir) (elementAssembler, error) {
	n := countDolbyVision(d)
	if n == 0 {
		return p.defaultElementAssembler, nil
	}
	if n != len(d.File) {
		return nil, ErrMixedPresets
	}
	_ = p.dolbyVisionElementAssembler // switch back to this once Hybrik fixes bug with GCP jobs hanging
	return p.dolbyVisionLegacyElementAssembler, nil
}
