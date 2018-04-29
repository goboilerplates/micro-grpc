package util

import (
	"github.com/goboilerplates/core"
	"github.com/goboilerplates/micro-rpc/proto"
)

// ConvertCoreSampleToReplySample converts core sample into the reply sample.
func ConvertCoreSampleToReplySample(coreSample *core.SampleEntity) *proto.GetSamplesReply_Sample {
	return &proto.GetSamplesReply_Sample{
		Name: coreSample.Name,
	}
}

// ConvertCoreSampleArrayToReplySampleArray converts core sample array into the replay sample array.
func ConvertCoreSampleArrayToReplySampleArray(coreSamples []*core.SampleEntity) []*proto.GetSamplesReply_Sample {
	var result []*proto.GetSamplesReply_Sample
	for _, item := range coreSamples {
		sample := ConvertCoreSampleToReplySample(item)
		result = append(result, sample)
	}
	return result
}
