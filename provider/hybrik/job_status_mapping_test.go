package hybrik

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	hy "github.com/cbsinteractive/hybrik-sdk-go"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func TestFiles(t *testing.T) {
	type Out = job.File

	tests := []struct {
		name, file           string
		want                 []job.File
		expectMissingOutputs bool
	}{
		{
			name: "task_status_combine_segments.json",
			file: "testdata/task_status_combine_segments.json",
			want: []Out{
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/046afa431fe57178/CBS_NCISNO_509_AA_30M_6CH_528.mp4",
					Container: "mp4",
					Size:      163718122,
				},
			},
		},
		{
			name: "task_status_dovi_transcode.json",
			file: "testdata/task_status_dovi_transcode.json",
			want: []Out{
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/blackmonday_540.mp4",
					Container: "mp4",
					Size:      492007718,
				},
			},
		},
		{
			name: "task_status_filename_no_extension.json",
			file: "testdata/task_status_filename_no_extension.json",
			want: []Out{
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/046afa431fe57178/CBS_NCISNO_509_AA_30M_6CH_528",
					Container: "mp4",
					Size:      163718122,
				},
			},
		},
		{
			name: "task_status_legacy_dovi_post_process.json",
			file: "testdata/task_status_legacy_dovi_post_process.json",
			want: []Out{
				{
					Name:      "gs://mediahub-dev/encodes/old_structure/733cc64ccde05511/dovi_custom_filename_1.mp4",
					Container: "mp4",
					Size:      6998510,
				},
			},
		},
		{
			name: "task_status_package.json",
			file: "testdata/task_status_package.json",
			want: []Out{
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/master.m3u8",
					Container: "m3u8",
					Size:      1429,
				},
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_360_audio.m3u8",
					Container: "m3u8",
					Size:      16328,
				},
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_360_video.m3u8",
					Container: "m3u8",
					Size:      15402,
				},
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_540_video.m3u8",
					Container: "m3u8",
					Size:      15402,
				},
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_540_audio.m3u8",
					Container: "m3u8",
					Size:      16328,
				},
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_540-iframes.m3u8",
					Container: "m3u8",
					Size:      22646,
				},
				{
					Name:      "s3://vtg-tsymborski-test-bucket/encodes/blackmonday/hls/blackmonday_360-iframes.m3u8",
					Container: "m3u8",
					Size:      22628,
				},
			},
		},
		{
			name:                 "task_status_no_outputs.json",
			file:                 "testdata/task_status_no_outputs.json",
			expectMissingOutputs: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var taskResult hy.TaskResult
			file, _ := ioutil.ReadFile(tt.file)
			err := json.Unmarshal(file, &taskResult)
			if err != nil {
				t.Fatal(err)
			}

			files, found, err := filesFrom(taskResult)
			if err != nil {
				t.Fatal(err)
			}

			if found && tt.expectMissingOutputs {
				t.Fatal("expected no outputs to be found")
			}
			for i, w := range files {
				h := files[i]
				if !reflect.DeepEqual(h, w) {
					t.Fatalf("file[%d]:\n\t\thave: %v\n\t\twant: %v", i, h, w)
				}
			}
			if h, w := len(file), len(tt.want); h != w {
				t.Errorf("bad file count: have %d want %d", h, w)
			}
		})
	}
}
