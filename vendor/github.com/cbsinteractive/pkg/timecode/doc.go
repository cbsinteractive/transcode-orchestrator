// Package timecode deals with timecode intervals, useful for
// operations on media timestamps. The two primary types in
// this package are:
//
// 	type Range [2]float64
//
// 	and
//
// 	type Splice []Range
//
// The Range represents a time interval expressed in seconds
// and the Splice is a collection of (possibly ordered, possibly
// overlapping ranges).
//
// The Splice type implments several utility methods to compute
// input and output timecodes. It also implements sort.Interface
// to assist with ordering timecodes
package timecode
