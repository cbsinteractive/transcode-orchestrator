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
	Size      int64         `json:"size,omitempty"`
	Duration  time.Duration `json:"dur,omitempty"`
	Name      string        `json:"name,omitempty"`
	Container string        `json:"container,omitempty"`
	Video     Video         `json:"video,omitempty"`
	Audio     Audio         `json:"audio,omitempty"`

	Splice                  timecode.Splice `json:"splice,omitempty"`
	Downmix                 *Downmix
	ExplicitKeyframeOffsets []float64
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
