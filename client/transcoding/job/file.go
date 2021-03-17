package job

import (
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/cbsinteractive/pkg/timecode"
)

// File
type File struct {
	// Name is the name of the file, which can be a full URI
	// absolute, relative path, or any meaningful string
	Name string `json:"name,omitempty"`

	Size  int64 `json:"size,omitempty"`
	Video Video `json:"video,omitempty"`
	Audio Audio `json:"audio,omitempty"`

	Duration time.Duration `json:"dur,omitempty"`
	// Splicing specifies ranges to chop and concatenate
	Splice timecode.Splice `json:"splice,omitempty"`

	// Downmix is kind of a hack here, it contains a set
	// of audio channels for one input File and one or more
	// output Files.
	//
	// We should refactor this to put the audio channel info
	// under the Audio struct, and then compute the Mix
	// by applying the channels from the source to the
	// desired destination files.
	//
	// As of now, the Downmix object is set on the source File
	// as a convention until we can do this part of the refactor
	Downmix *Downmix `json:"downmix,omitempty"`

	// ExplicitKeyframeOffsets should probably be in the video section
	ExplicitKeyframeOffsets []float64 `json:",omitempty"`

	// NOTE(as): I *really* hope we can deprecate this. Our
	// mediahub code path makes it impossible for this to
	// differ from the file's extension, but for some reason
	// we have a test that specifically requests this impossible
	// condition
	Container string `json:"container,omitempty"`
}

func (f File) Join(name string) File {
	u := f.URL()
	u.Path = path.Join(u.Path, name)
	f.Name = u.String()
	return f
}

func (f File) URL() url.URL {
	u, _ := url.Parse(f.Name)
	if u == nil {
		return url.URL{}
	}
	return *u
}
func (f File) Provider() string {
	return f.URL().Scheme
}
func (f File) Type() string {
	return strings.TrimPrefix(path.Ext(f.URL().Path), ".")
}
func (f File) Dir() string {
	if f.Type() == "" {
		return f.Name
	}
	u := f.URL()
	u.Path = path.Dir(u.Path)
	return u.String()
}
func (f File) Base() string {
	return path.Base(f.URL().Path)
}

// Dir is a named directory of files
type Dir struct {
	Path string `json:"path,omitempty"`
	File []File `json:"files,omitempty"`
}

func (d *Dir) Len() int {
	return len(d.File)
}

func (d *Dir) Add(f ...File) {
	d.File = append(d.File, f...)
}

func (d Dir) Location() url.URL {
	u, _ := url.Parse(d.Path)
	if u == nil {
		return url.URL{}
	}
	return *u
}
