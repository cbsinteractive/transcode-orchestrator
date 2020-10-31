package bitmovin

import (
	"context"
	"fmt"
	"sort"

	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/pkg/timecode"
)

func (p *bitmovinProvider) trace(ctx context.Context, name string, err *error) func() {
	x := p.tracer.BeginSubsegment(ctx, name)
	return func() {
		if err == nil {
			x.Close(nil)
		} else {
			x.Close(*err)
		}
	}
}

func (p *bitmovinProvider) splice(ctx context.Context, encID, stream0 string, splice timecode.Splice) (stream1 string, err error) {
	defer p.trace(ctx, "bitmovin-create-concatenated-splice", &err)()

	if len(splice) == 0 {
		return stream0, nil
	}

	type work struct {
		pos        int32
		start, dur float64
		id         string
		err        error
	}
	workc := make(chan work, len(splice))

	// splice each range concurrently
	for i, r := range splice {
		w := work{
			pos:   int32(i),
			start: r[0],
			dur:   r[1] - r[0],
		}
		go func() {
			// NOTE(as): don't use the timecode "api", it seems to look for a real
			// timecode track in the source. If it doesn't find it, it just doesn't trim
			// the clip and provides no logging or errors. For this "api", it wants
			// start, duration; not start, end, and it also wants pointers
			splice, err := p.api.Encoding.Encodings.InputStreams.Trimming.TimeBased.Create(encID, model.TimeBasedTrimmingInputStream{
				InputStreamId: stream0,
				Offset:        &w.start,
				Duration:      &w.dur,
			})
			if splice != nil {
				w.id = splice.Id
			}
			w.err = err
			workc <- w
		}()
	}

	if cap(workc) == 1 {
		// NOTE(as): turns out bitmovin complains if you run the equivalent of:
		// 'cat input0.mp4 > input.mp4'  because there's only one input0.mp4
		w := <-workc
		if w.err != nil {
			return stream0, fmt.Errorf("trim: range#%d: %w", 1, w.err)
		}
		// can't concatenate, need special case for one input splice
		return w.id, nil
	}

	// collect the results serially
	cat := []model.ConcatenationInputConfiguration{}
	for i := 0; i < cap(workc); i++ {
		w := <-workc
		if w.err != nil {
			return stream0, fmt.Errorf("trim: range#%d: %w", w.pos, w.err)
		}
		main := i == 0
		cat = append(cat, model.ConcatenationInputConfiguration{
			IsMain:        &main,
			InputStreamId: w.id,
			Position:      &w.pos,
		})
	}

	// although there are position markers in the struct,
	// sort it just in case, this makes the logging consistent too
	sort.Slice(cat, func(i, j int) bool {
		return *cat[i].Position < *cat[j].Position
	})
	c, err := p.api.Encoding.Encodings.InputStreams.Concatenation.Create(encID, model.ConcatenationInputStream{
		Concatenation: cat,
	})
	if err != nil {
		return stream0, fmt.Errorf("concatenation: %v", err)
	}
	return c.Id, nil
}
