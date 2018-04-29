package util_test

import (
	"testing"

	"github.com/goboilerplates/core"
	"github.com/stretchr/testify/assert"

	"github.com/goboilerplates/micro-rpc/util"
)

// SetupSample .
func SetupSample() core.SampleEntity {
	return core.SampleEntity{Name: "Kaka"}
}

// TestConvertCoreSampleToReplySample .
func TestConvertCoreSampleToReplySample(test *testing.T) {
	sample := SetupSample()
	result := util.ConvertCoreSampleToReplySample(&sample)

	assert.NotNil(test, result)
}

// TestConvertCoreSampleArrayToReplySampleArray .
func TestConvertCoreSampleArrayToReplySampleArray(test *testing.T) {
	sample := SetupSample()
	sampleArray := []*core.SampleEntity{&sample}
	result := util.ConvertCoreSampleArrayToReplySampleArray(sampleArray)

	assert.NotNil(test, result)
}
