package provider

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cbsinteractive/transcode-orchestrator/config"
)

func TestListProviders(t *testing.T) {
	cap := Capabilities{
		InputFormats:  []string{"prores", "h264"},
		OutputFormats: []string{"mp4", "hls"},
		Destinations:  []string{"s3", "akamai"},
	}
	providers = map[string]Factory{
		"cap-and-unhealthy": getFactory(nil, errors.New("api is down"), cap),
		"factory-err":       getFactory(errors.New("invalid config"), nil, cap),
		"cap-and-healthy":   getFactory(nil, nil, cap),
	}
	expected := []string{"cap-and-healthy", "cap-and-unhealthy"}
	got := List(&config.Config{})
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("DescribeProviders: want %#v. Got %#v", expected, got)
	}
}

func TestListProvidersEmpty(t *testing.T) {
	providers = nil
	providerNames := List(&config.Config{})
	if len(providerNames) != 0 {
		t.Errorf("Unexpected non-empty provider list: %#v", providerNames)
	}
}

func TestDescribe(t *testing.T) {
	cap := Capabilities{
		InputFormats:  []string{"prores", "h264"},
		OutputFormats: []string{"mp4", "hls"},
		Destinations:  []string{"s3", "akamai"},
	}
	providers = map[string]Factory{
		"cap-and-unhealthy": getFactory(nil, errors.New("api is down"), cap),
		"factory-err":       getFactory(errors.New("invalid config"), nil, cap),
		"cap-and-healthy":   getFactory(nil, nil, cap),
	}
	var tests = []struct {
		input    string
		expected Description
	}{
		{
			"factory-err",
			Description{Name: "factory-err", Enabled: false},
		},
		{
			"cap-and-healthy",
			Description{
				Name:         "cap-and-healthy",
				Capabilities: cap,
				Health:       Health{OK: true},
				Enabled:      true,
			},
		},
		{
			"cap-and-unhealthy",
			Description{
				Name:         "cap-and-unhealthy",
				Capabilities: cap,
				Health:       Health{OK: false, Message: "api is down"},
				Enabled:      true,
			},
		},
	}
	for _, test := range tests {
		desc, err := Describe(test.input, &config.Config{})
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(*desc, test.expected) {
			t.Errorf("DescribeProvider(%q): want %#v. Got %#v", test.input, test.expected, *desc)
		}
	}

	description, err := Describe("anything", nil)
	if err != ErrNotFound {
		t.Errorf("Wrong error. Want %#v. Got %#v", ErrNotFound, err)
	}
	if description != nil {
		t.Errorf("Unexpected non-nil description: %#v", description)
	}
}
