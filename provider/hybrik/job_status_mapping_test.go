package hybrik

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/google/go-cmp/cmp"
)

func Test_filesFrom(t *testing.T) {
	tests := []struct {
		name, file           string
		outputFiles          []provider.OutputFile
		expectMissingOutputs bool
	}{
		{
			name: "pulls the correct data from combine segments tasks",
			file: "testdata/task_status_combine_segments.json",
			outputFiles: []provider.OutputFile{
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/046afa431fe57178/CBS_NCISNO_509_AA_30M_6CH_528.mp4",
					Container: "mp4",
					FileSize:  163718122,
				},
			},
		},
		{
			name: "pulls the correct data from dolby vision transcode tasks",
			file: "testdata/task_status_dovi_transcode.json",
			outputFiles: []provider.OutputFile{
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/blackmonday_540.mp4",
					Container: "mp4",
					FileSize:  492007718,
				},
			},
		},
		{
			name: "returns a valid container when the filename contains no extension",
			file: "testdata/task_status_filename_no_extension.json",
			outputFiles: []provider.OutputFile{
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/046afa431fe57178/CBS_NCISNO_509_AA_30M_6CH_528",
					Container: "mp4",
					FileSize:  163718122,
				},
			},
		},
		{
			name: "pulls the correct data from legacy dolby vision post-process tasks",
			file: "testdata/task_status_legacy_dovi_post_process.json",
			outputFiles: []provider.OutputFile{
				{
					Path:      "gs://mediahub-dev/encodes/old_structure/733cc64ccde05511/dovi_custom_filename_1.mp4",
					Container: "mp4",
					FileSize:  6998510,
				},
			},
		},
		{
			name: "pulls the correct data from package tasks",
			file: "testdata/task_status_package.json",
			outputFiles: []provider.OutputFile{
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/master.m3u8",
					Container: "m3u8",
					FileSize:  1429,
				},
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_360_audio.m3u8",
					Container: "m3u8",
					FileSize:  16328,
				},
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_360_video.m3u8",
					Container: "m3u8",
					FileSize:  15402,
				},
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_540_video.m3u8",
					Container: "m3u8",
					FileSize:  15402,
				},
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_540_audio.m3u8",
					Container: "m3u8",
					FileSize:  16328,
				},
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_540-iframes.m3u8",
					Container: "m3u8",
					FileSize:  22646,
				},
				{
					Path:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_360-iframes.m3u8",
					Container: "m3u8",
					FileSize:  22628,
				},
			},
		},
		{
			name:                 "does not find outputs in files that are not recognized as containing them",
			file:                 "testdata/task_status_no_outputs.json",
			expectMissingOutputs: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			file, err := ioutil.ReadFile(tt.file)
			if err != nil {
				t.Error(err)
				return
			}

			var taskResult hybrik.TaskResult
			err = json.Unmarshal(file, &taskResult)
			if err != nil {
				t.Error(err)
				return
			}

			files, found, err := filesFrom(taskResult)
			if err != nil {
				t.Error(err)
				return
			}

			if found && tt.expectMissingOutputs {
				t.Errorf("expected no outputs to be found")
				return
			}

			if g, e := files, tt.outputFiles; !reflect.DeepEqual(g, e) {
				t.Errorf("wrong jobs: got %v\nexpected %v\ndiff %v", g, e, cmp.Diff(g, e))
			}
		})
	}
}
