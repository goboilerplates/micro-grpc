package util_test

import (
	"testing"

	"github.com/goboilerplates/micro-rpc/example/goclient/util"
	"github.com/stretchr/testify/assert"
)

// TestGetCurrentTimeStampe .
func TestGetCurrentTimeStampe(test *testing.T) {
	result := util.GetCurrentTimeStamp()

	assert.NotNil(test, result)
}
