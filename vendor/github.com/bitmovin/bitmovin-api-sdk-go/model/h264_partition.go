package model
// H264Partition : Partitions to consider. Analyzing more partition options improves quality at the cost of speed.
type H264Partition string

// List of H264Partition
const (
	H264Partition_NONE H264Partition = "NONE"
	H264Partition_P8X8 H264Partition = "P8X8"
	H264Partition_P4X4 H264Partition = "P4X4"
	H264Partition_B8X8 H264Partition = "B8X8"
	H264Partition_I8X8 H264Partition = "I8X8"
	H264Partition_I4X4 H264Partition = "I4X4"
	H264Partition_ALL H264Partition = "ALL"
)