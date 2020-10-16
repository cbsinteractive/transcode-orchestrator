package container

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/test"
	"github.com/google/go-cmp/cmp"
)

func TestCMAFAssembler(t *testing.T) {
	defaultAssemblerCfg := AssemblerCfg{
		EncID:    "testEncID",
		OutputID: "testOutputID",
		AudCfgID: "testAudCfgID",
		VidCfgID: "testVidCfgID",
		AudMuxingStream: model.MuxingStream{
			StreamId: "testAudStreamID",
		},
		VidMuxingStream: model.MuxingStream{
			StreamId: "testVidStreamID",
		},
		ManifestID:         "testManifestID",
		ManifestMasterPath: "test/master/manifest/path",
		SegDuration:        88,
	}

	tests := []struct {
		name         string
		cfg          AssemblerCfg
		api          CMAFContainerAPI
		wantErr      string
		assertParams func(*testing.T, CMAFContainerAPI)
	}{
		{
			name: "an hls config with video and audio result in the correct calls to the HLSContainerAPI",
			cfg:  defaultAssemblerCfg,
			api:  cmafContainerAPI(),
			assertParams: func(t *testing.T, api CMAFContainerAPI) {
				cmafMuxingAPI := api.CMAFMuxing.(*fakeCMAFMuxingAPI)
				hlsStreamsAPI := api.HLSStreams.(*fakeHLSStreamsAPI)
				hlsAudioMediaAPI := api.HLSAudioMedia.(*fakeHLSAudioMediaAPI)

				if g, e := cmafMuxingAPI.numInvocations, 2; g != e {
					t.Errorf("invalid number of calls to the TS muxing api: got %d, expected %d", g, e)
					return
				}

				if g, e := cmafMuxingAPI.invocationDetails[0].encodingID, "testEncID"; g != e {
					t.Errorf("invalid encodingID sent: got %q, expected %q", g, e)
				}

				expectedAudioMuxing := model.CmafMuxing{
					Streams: []model.MuxingStream{{StreamId: "testAudStreamID"}},
					Outputs: []model.EncodingOutput{{
						OutputId:   "testOutputID",
						OutputPath: "test/master/manifest/path/testAudCfgID",
						Acl:        []model.AclEntry{{Permission: "PRIVATE"}},
					}},
					SegmentLength: floatToPtr(88),
					SegmentNaming: "seg_%number%.m4a",
				}

				if g, e := cmafMuxingAPI.invocationDetails[0].muxing, expectedAudioMuxing; !reflect.DeepEqual(g, e) {
					t.Errorf("invalid audio muxing: got  %v\nexpected %v\ndiff %v", g, e, cmp.Diff(g, e))
				}

				expectedVideoMuxing := model.CmafMuxing{
					Streams: []model.MuxingStream{{StreamId: "testVidStreamID"}},
					Outputs: []model.EncodingOutput{{
						OutputId:   "testOutputID",
						OutputPath: "test/master/manifest/path/testVidCfgID",
						Acl:        []model.AclEntry{{Permission: "PRIVATE"}},
					}},
					SegmentLength: floatToPtr(88),
					SegmentNaming: "seg_%number%.m4v",
				}

				if g, e := cmafMuxingAPI.invocationDetails[1].muxing, expectedVideoMuxing; !reflect.DeepEqual(g, e) {
					t.Errorf("invalid video muxing: got  %v\nexpected %v\ndiff %v", g, e, cmp.Diff(g, e))
				}

				// HLSStreamsAPI tests

				if g, e := hlsStreamsAPI.numInvocations, 1; g != e {
					t.Errorf("invalid number of calls to the HLS streams api: got %d, expected %d", g, e)
					return
				}

				if g, e := hlsStreamsAPI.invocationDetails[0].manifestID, "testManifestID"; g != e {
					t.Errorf("invalid manifestID sent: got %q, expected %q", g, e)
				}

				expectedHLSStreamInfo := model.StreamInfo{
					Audio:       "testAudCfgID",
					EncodingId:  "testEncID",
					StreamId:    "testVidStreamID",
					SegmentPath: "test/master/manifest/path/testVidCfgID",
					Uri:         "testVidCfgID.m3u8",
				}

				if g, e := hlsStreamsAPI.invocationDetails[0].streamInfo, expectedHLSStreamInfo; !reflect.DeepEqual(g, e) {
					t.Errorf("invalid hls stream info: got  %v\nexpected %v\ndiff %v", g, e, cmp.Diff(g, e))
				}

				// HLSAudioMediaAPI tests

				if g, e := hlsAudioMediaAPI.numInvocations, 1; g != e {
					t.Errorf("invalid number of calls to the HLS audio media api: got %d, expected %d", g, e)
					return
				}

				if g, e := hlsAudioMediaAPI.invocationDetails[0].manifestID, "testManifestID"; g != e {
					t.Errorf("invalid manifestID sent: got %q, expected %q", g, e)
				}

				expectedHLSAudioMedia := model.AudioMediaInfo{
					GroupId:         "testAudCfgID",
					Language:        "en",
					Name:            "testAudCfgID",
					IsDefault:       boolToPtr(false),
					Autoselect:      boolToPtr(false),
					Characteristics: []string{"public.accessibility.describes-video"},
					SegmentPath:     "testAudCfgID",
					EncodingId:      "testEncID",
					StreamId:        "testAudStreamID",
					Uri:             "testAudCfgID.m3u8",
					Forced:          boolToPtr(false),
				}

				if g, e := hlsAudioMediaAPI.invocationDetails[0].mediaInfo, expectedHLSAudioMedia; !reflect.DeepEqual(g, e) {
					t.Errorf("invalid hls audio media: got  %v\nexpected %v\ndiff %v", g, e, cmp.Diff(g, e))
				}
			},
		},
		{
			name: "an hls config with only video results in the correct calls to the HLSContainerAPI",
			cfg: AssemblerCfg{
				EncID:    "testEncID",
				OutputID: "testOutputID",
				VidCfgID: "testVidCfgID",
				VidMuxingStream: model.MuxingStream{
					StreamId: "testVidStreamID",
				},
				ManifestID:         "testManifestID",
				ManifestMasterPath: "test/master/manifest/path",
				SegDuration:        88,
			},
			api: cmafContainerAPI(),
			assertParams: func(t *testing.T, api CMAFContainerAPI) {
				cmafMuxingAPI := api.CMAFMuxing.(*fakeCMAFMuxingAPI)
				hlsStreamsAPI := api.HLSStreams.(*fakeHLSStreamsAPI)
				hlsAudioMediaAPI := api.HLSAudioMedia.(*fakeHLSAudioMediaAPI)

				if g, e := cmafMuxingAPI.numInvocations, 1; g != e {
					t.Errorf("invalid number of calls to the TS muxing api: got %d, expected %d", g, e)
					return
				}

				if g, e := cmafMuxingAPI.invocationDetails[0].encodingID, "testEncID"; g != e {
					t.Errorf("invalid encodingID sent: got %q, expected %q", g, e)
				}

				expectedVideoMuxing := model.CmafMuxing{
					Streams: []model.MuxingStream{{StreamId: "testVidStreamID"}},
					Outputs: []model.EncodingOutput{{
						OutputId:   "testOutputID",
						OutputPath: "test/master/manifest/path/testVidCfgID",
						Acl:        []model.AclEntry{{Permission: "PRIVATE"}},
					}},
					SegmentLength: floatToPtr(88),
					SegmentNaming: "seg_%number%.m4v",
				}

				if g, e := cmafMuxingAPI.invocationDetails[0].muxing, expectedVideoMuxing; !reflect.DeepEqual(g, e) {
					t.Errorf("invalid video muxing: got  %v\nexpected %v\ndiff %v", g, e, cmp.Diff(g, e))
				}

				// HLSStreamsAPI tests

				if g, e := hlsStreamsAPI.numInvocations, 1; g != e {
					t.Errorf("invalid number of calls to the HLS streams api: got %d, expected %d", g, e)
					return
				}

				if g, e := hlsStreamsAPI.invocationDetails[0].manifestID, "testManifestID"; g != e {
					t.Errorf("invalid manifestID sent: got %q, expected %q", g, e)
				}

				expectedHLSStreamInfo := model.StreamInfo{
					EncodingId:  "testEncID",
					StreamId:    "testVidStreamID",
					SegmentPath: "test/master/manifest/path/testVidCfgID",
					Uri:         "testVidCfgID.m3u8",
				}

				if g, e := hlsStreamsAPI.invocationDetails[0].streamInfo, expectedHLSStreamInfo; !reflect.DeepEqual(g, e) {
					t.Errorf("invalid hls stream info: got  %v\nexpected %v\ndiff %v", g, e, cmp.Diff(g, e))
				}

				// HLSAudioMediaAPI tests

				if g, e := hlsAudioMediaAPI.numInvocations, 0; g != e {
					t.Errorf("invalid number of calls to the HLS audio media api: got %d, expected %d", g, e)
					return
				}
			},
		},
		{
			name: "when the ts muxing api is erroring, a useful error is returned",
			cfg:  defaultAssemblerCfg,
			api: CMAFContainerAPI{
				HLSAudioMedia: &fakeHLSAudioMediaAPI{},
				CMAFMuxing:    &fakeCMAFMuxingAPI{forceErr: true},
				HLSStreams:    &fakeHLSStreamsAPI{},
			},
			wantErr: "creating audio cmaf muxing: forced by test",
		},
		{
			name: "when the hls audio media api is erroring, a useful error is returned",
			cfg:  defaultAssemblerCfg,
			api: CMAFContainerAPI{
				HLSAudioMedia: &fakeHLSAudioMediaAPI{forceErr: true},
				CMAFMuxing:    &fakeCMAFMuxingAPI{},
				HLSStreams:    &fakeHLSStreamsAPI{},
			},
			wantErr: "creating audio media: forced by test",
		},
		{
			name: "when the hls streams api is erroring, a useful error is returned",
			cfg:  defaultAssemblerCfg,
			api: CMAFContainerAPI{
				HLSAudioMedia: &fakeHLSAudioMediaAPI{},
				CMAFMuxing:    &fakeCMAFMuxingAPI{},
				HLSStreams:    &fakeHLSStreamsAPI{forceErr: true},
			},
			wantErr: "creating video stream info: forced by test",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			err := NewCMAFAssembler(tt.api).Assemble(tt.cfg)
			if shouldReturn := test.AssertWantErr(err, tt.wantErr, "Assemble()", t); shouldReturn {
				return
			}

			if paramsAssertion := tt.assertParams; paramsAssertion != nil {
				paramsAssertion(t, tt.api)
			}
		})
	}
}

func cmafContainerAPI() CMAFContainerAPI {
	return CMAFContainerAPI{
		HLSAudioMedia: &fakeHLSAudioMediaAPI{},
		CMAFMuxing:    &fakeCMAFMuxingAPI{},
		HLSStreams:    &fakeHLSStreamsAPI{},
	}
}

type fakeCMAFMuxingAPI struct {
	forceErr          bool
	numInvocations    int
	invocationDetails []struct {
		encodingID string
		muxing     model.CmafMuxing
	}
}

func (a *fakeCMAFMuxingAPI) Create(encodingID string, cmafMuxing model.CmafMuxing) (*model.CmafMuxing, error) {
	if a.forceErr {
		return nil, errors.New("forced by test")
	}
	a.numInvocations++
	a.invocationDetails = append(a.invocationDetails, struct {
		encodingID string
		muxing     model.CmafMuxing
	}{encodingID: encodingID, muxing: cmafMuxing})

	return &cmafMuxing, nil
}
