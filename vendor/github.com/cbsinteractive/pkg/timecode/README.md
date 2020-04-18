## 

```
	// A range is a closed pair of durations representing
	// a start and end time. A splice is a list of ranges.

	// ParseTimecode can decode timestamps, with an
	// optional non-zero framerate argument (60). It returns
	// a Range. In this case we have a 10 second video at
	// 60 frames-per-second.
	r, _ := ParseTimecode("00:00:10:00", 60)
	fmt.Println("video timecode", r.Timecode(60))

	// Suppose we want to extract the footage between
	// 1s-5s. After we define this with a splice, we can
	// check to see if our splice is in the Range.
	s := Splice{{1,5}}
	fmt.Println(s, "in", r, "?", s.In(r))

	// We can also check if this splice is sorted. Sorting
	// in a splice is done by start time and range duration
	// in that order.
	fmt.Println("splice sorted?", s.Sorted())

	// If we wanted to process this 10 second clip, we
	// wouldn't need to send the whole clip to the provider.
	// Union can tell us what section of video we would need
	// to transmit, saving bandwidth.
	fmt.Println("input: minimum span", s.Union())

	// Finally, we can compute the output duration of the
	// final clip.
	fmt.Println("output length", s.Size())
```
	