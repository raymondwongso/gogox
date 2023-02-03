package log

import "golang.org/x/exp/maps"

// Metadata defines metadata for logger, which will be included whenever log is written.
type Metadata map[string]interface{}

// MergeMetadata creates new metadata as a result from merging between md1 and md2.
func MergeMetadata(md1, md2 Metadata) Metadata {
	res := Metadata{}
	maps.Copy(res, md1)
	maps.Copy(res, md2)
	return res
}
