package util

import (
	"github.com/goboilerplates/core"
	"github.com/goboilerplates/micro-rpc/proto"
)

// ConvertCoreSampleToReplySample .
func ConvertCoreSampleToReplySample(coreSample *core.SampleEntity) *proto.GetSamplesReply_Sample {
	return &proto.GetSamplesReply_Sample{
		Name: coreSample.Name,
	}
}

// ConvertCoreSampleArrayToReplySampleArray .
func ConvertCoreSampleArrayToReplySampleArray(coreSamples []*core.SampleEntity) []*proto.GetSamplesReply_Sample {
	var result []*proto.GetSamplesReply_Sample
	for _, item := range coreSamples {
		sample := ConvertCoreSampleToReplySample(item)
		result = append(result, sample)
	}
	return result
}
