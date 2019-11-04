package hybrik

import (
	"fmt"
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/video-transcoding-api/provider"
)

type taskWithOutputMatcher struct {
	kind     string
	uidRegex *regexp.Regexp
}

var tasksWithOutputsMatchers []taskWithOutputMatcher

func init() {
	doViPostProcessRegex, err := regexp.Compile(`post_transcode_stage_[\d]+$`)
	if err != nil {
		log.Panicf("compiling the doVi post process regex: %v", err)
	}

	doViTranscodeRegex, err := regexp.Compile(`dolby_vision_[\d]+$`)
	if err != nil {
		log.Panicf("compiling the doVi transcode regex: %v", err)
	}

	transcodeRegex, err := regexp.Compile(`transcode_task_[\d]+$`)
	if err != nil {
		log.Panicf("compiling the transcode regex: %v", err)
	}

	packageRegex, err := regexp.Compile(`packager$`)
	if err != nil {
		log.Panicf("compiling the package regex: %v", err)
	}

	combinerRegex, err := regexp.Compile(`combiner_[\d]+$`)
	if err != nil {
		log.Panicf("compiling the combiner regex: %v", err)
	}

	tasksWithOutputsMatchers = []taskWithOutputMatcher{
		{kind: "Dolby Vision", uidRegex: doViPostProcessRegex},
		{kind: "Dolby Vision", uidRegex: doViTranscodeRegex},
		{kind: "Transcode", uidRegex: transcodeRegex},
		{kind: "Package", uidRegex: packageRegex},
		{kind: "Combine Segments", uidRegex: combinerRegex},
	}
}

func filesFrom(task hybrik.TaskResult) ([]provider.OutputFile, bool, error) {
	// ensure the task type results in outputs
	if !taskHasOutputs(task, tasksWithOutputsMatchers) {
		return nil, false, nil
	}

	var files []provider.OutputFile
	for _, document := range task.Documents {
		for _, assetVersion := range document.ResultPayload.Payload.AssetVersions {
			for _, component := range assetVersion.AssetComponents {
				normalizedPath := strings.TrimRight(assetVersion.Location.Path, "/")
				files = append(files, provider.OutputFile{
					Path:      fmt.Sprintf("%s/%s", normalizedPath, component.Name),
					Container: containerFrom(component),
					FileSize:  int64(component.Descriptor.Size),
				})
			}
		}
	}

	return files, len(files) > 0, nil
}

const assetMediaInfoType = "ASSET"

func containerFrom(component hybrik.AssetComponentResult) string {
	if infos := component.MediaAnalyze.MediaInfo; len(infos) > 0 {
		for _, i := range infos {
			if i.StreamType == assetMediaInfoType && i.ASSET.Format != "" {
				return i.ASSET.Format
			}
		}
	}

	return strings.Replace(path.Ext(component.Name), ".", "", -1)
}

func taskHasOutputs(task hybrik.TaskResult, matchers []taskWithOutputMatcher) bool {
	for _, matcher := range matchers {
		if matcher.kind != task.Kind {
			continue
		}

		if matcher.uidRegex.Match([]byte(task.UID)) {
			return true
		}
	}

	return false
}
