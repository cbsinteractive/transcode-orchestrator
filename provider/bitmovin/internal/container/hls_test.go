package container

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/cbsinteractive/video-transcoding-api/test"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

func TestNewHLSAssembler(t *testing.T) {
	defaultAssemblerCfg := AssemblerCfg{
		EncID:       "testEncID",
		OutputID:    "testOutputID",
		AudCfgID:    "testAudCfgID",
		VidCfgID:    "testVidCfgID",
		AudStreamID: "testAudStreamID",
		VidStreamID: "testVidStreamID",
		AudMuxingStream: model.MuxingStream{
			StreamId: "testAudMuxingStreamID",
		},
		VidMuxingStream: model.MuxingStream{
			StreamId: "testVidMuxingStreamID",
		},
		ManifestID:         "testManifestID",
		ManifestMasterPath: "test/master/manifest/path",
		SegDuration:        88,
	}

	tests := []struct {
		name         string
		cfg          AssemblerCfg
		api          HLSContainerAPI
		wantErr      string
		assertParams func(*testing.T, HLSContainerAPI)
	}{
		{
			name: "an hls config with video and audio result in the correct calls to the HLSContainerAPI",
			cfg:  defaultAssemblerCfg,
			api:  hlsContainerAPI(),
			assertParams: func(t *testing.T, api HLSContainerAPI) {
				tsMuxingAPI := api.TSMuxing.(*fakeTSMuxingAPI)
				hlsStreamsAPI := api.HLSStreams.(*fakeHLSStreamsAPI)
				hlsAudioMediaAPI := api.HLSAudioMedia.(*fakeHLSAudioMediaAPI)

				if g, e := tsMuxingAPI.numInvocations, 2; g != e {
					t.Errorf("invalid number of calls to the TS muxing api: got %d, expected %d", g, e)
					return
				}

				if g, e := tsMuxingAPI.invocationDetails[0].encodingID, "testEncID"; g != e {
					t.Errorf("invalid encodingID sent: got %q, expected %q", g, e)
				}

				expectedAudioMuxing := model.TsMuxing{
					Streams: []model.MuxingStream{{StreamId: "testAudMuxingStreamID"}},
					Outputs: []model.EncodingOutput{{
						OutputId:   "testOutputID",
						OutputPath: "test/master/manifest/path/testAudCfgID",
						Acl:        []model.AclEntry{{Permission: "PRIVATE"}},
					}},
					SegmentLength: floatToPtr(88),
					SegmentNaming: "seg_%number%.ts",
				}

				if g, e := tsMuxingAPI.invocationDetails[0].muxing, expectedAudioMuxing; !reflect.DeepEqual(g, e) {
					t.Errorf("invalid audio muxing: got  %v\nexpected %v\ndiff %v", g, e, cmp.Diff(g, e))
				}

				expectedVideoMuxing := model.TsMuxing{
					Streams: []model.MuxingStream{{StreamId: "testVidMuxingStreamID"}},
					Outputs: []model.EncodingOutput{{
						OutputId:   "testOutputID",
						OutputPath: "test/master/manifest/path/testVidCfgID",
						Acl:        []model.AclEntry{{Permission: "PRIVATE"}},
					}},
					SegmentLength: floatToPtr(88),
					SegmentNaming: "seg_%number%.ts",
				}

				if g, e := tsMuxingAPI.invocationDetails[1].muxing, expectedVideoMuxing; !reflect.DeepEqual(g, e) {
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
					SegmentPath: "testVidCfgID",
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
				EncID:       "testEncID",
				OutputID:    "testOutputID",
				VidCfgID:    "testVidCfgID",
				VidStreamID: "testVidStreamID",
				VidMuxingStream: model.MuxingStream{
					StreamId: "testVidMuxingStreamID",
				},
				ManifestID:         "testManifestID",
				ManifestMasterPath: "test/master/manifest/path",
				SkipAudioCreation:  true,
				SegDuration:        88,
			},
			api: hlsContainerAPI(),
			assertParams: func(t *testing.T, api HLSContainerAPI) {
				tsMuxingAPI := api.TSMuxing.(*fakeTSMuxingAPI)
				hlsStreamsAPI := api.HLSStreams.(*fakeHLSStreamsAPI)
				hlsAudioMediaAPI := api.HLSAudioMedia.(*fakeHLSAudioMediaAPI)

				if g, e := tsMuxingAPI.numInvocations, 1; g != e {
					t.Errorf("invalid number of calls to the TS muxing api: got %d, expected %d", g, e)
					return
				}

				if g, e := tsMuxingAPI.invocationDetails[0].encodingID, "testEncID"; g != e {
					t.Errorf("invalid encodingID sent: got %q, expected %q", g, e)
				}

				expectedVideoMuxing := model.TsMuxing{
					Streams: []model.MuxingStream{{StreamId: "testVidMuxingStreamID"}},
					Outputs: []model.EncodingOutput{{
						OutputId:   "testOutputID",
						OutputPath: "test/master/manifest/path/testVidCfgID",
						Acl:        []model.AclEntry{{Permission: "PRIVATE"}},
					}},
					SegmentLength: floatToPtr(88),
					SegmentNaming: "seg_%number%.ts",
				}

				if g, e := tsMuxingAPI.invocationDetails[0].muxing, expectedVideoMuxing; !reflect.DeepEqual(g, e) {
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
					SegmentPath: "testVidCfgID",
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
			api: HLSContainerAPI{
				HLSAudioMedia: &fakeHLSAudioMediaAPI{},
				TSMuxing:      &fakeTSMuxingAPI{forceErr: true},
				HLSStreams:    &fakeHLSStreamsAPI{},
			},
			wantErr: "creating audio ts muxing: forced by test",
		},
		{
			name: "when the hls audio media api is erroring, a useful error is returned",
			cfg:  defaultAssemblerCfg,
			api: HLSContainerAPI{
				HLSAudioMedia: &fakeHLSAudioMediaAPI{forceErr: true},
				TSMuxing:      &fakeTSMuxingAPI{},
				HLSStreams:    &fakeHLSStreamsAPI{},
			},
			wantErr: "creating audio media: forced by test",
		},
		{
			name: "when the hls streams api is erroring, a useful error is returned",
			cfg:  defaultAssemblerCfg,
			api: HLSContainerAPI{
				HLSAudioMedia: &fakeHLSAudioMediaAPI{},
				TSMuxing:      &fakeTSMuxingAPI{},
				HLSStreams:    &fakeHLSStreamsAPI{forceErr: true},
			},
			wantErr: "creating video stream info: forced by test",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := NewHLSAssembler(tt.api).Assemble(tt.cfg)
			if shouldReturn := test.AssertWantErr(err, tt.wantErr, "Assemble()", t); shouldReturn {
				return
			}

			if paramsAssertion := tt.assertParams; paramsAssertion != nil {
				paramsAssertion(t, tt.api)
			}
		})
	}
}

func hlsContainerAPI() HLSContainerAPI {
	return HLSContainerAPI{
		HLSAudioMedia: &fakeHLSAudioMediaAPI{},
		TSMuxing:      &fakeTSMuxingAPI{},
		HLSStreams:    &fakeHLSStreamsAPI{},
	}
}

type fakeTSMuxingAPI struct {
	forceErr          bool
	numInvocations    int
	invocationDetails []struct {
		encodingID string
		muxing     model.TsMuxing
	}
}

func (a *fakeTSMuxingAPI) Create(encodingID string, tsMuxing model.TsMuxing) (*model.TsMuxing, error) {
	if a.forceErr {
		return nil, errors.New("forced by test")
	}
	a.numInvocations++
	a.invocationDetails = append(a.invocationDetails, struct {
		encodingID string
		muxing     model.TsMuxing
	}{encodingID: encodingID, muxing: tsMuxing})

	return &tsMuxing, nil
}

type fakeHLSAudioMediaAPI struct {
	forceErr          bool
	numInvocations    int
	invocationDetails []struct {
		manifestID string
		mediaInfo  model.AudioMediaInfo
	}
}

func (a *fakeHLSAudioMediaAPI) Create(manifestID string, mediaInfo model.AudioMediaInfo) (*model.AudioMediaInfo, error) {
	if a.forceErr {
		return nil, errors.New("forced by test")
	}
	a.numInvocations++
	a.invocationDetails = append(a.invocationDetails, struct {
		manifestID string
		mediaInfo  model.AudioMediaInfo
	}{manifestID, mediaInfo})

	return &mediaInfo, nil
}

type fakeHLSStreamsAPI struct {
	forceErr          bool
	numInvocations    int
	invocationDetails []struct {
		manifestID string
		streamInfo model.StreamInfo
	}
}

func (a *fakeHLSStreamsAPI) Create(manifestID string, streamInfo model.StreamInfo) (*model.StreamInfo, error) {
	if a.forceErr {
		return nil, errors.New("forced by test")
	}
	a.numInvocations++
	a.invocationDetails = append(a.invocationDetails, struct {
		manifestID string
		streamInfo model.StreamInfo
	}{manifestID, streamInfo})

	return &streamInfo, nil
}
