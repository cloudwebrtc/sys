package utilities_test

import (
	"github.com/getcouragenow/sys/sys-account/service/go/pkg/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenRandomByteSlice(t *testing.T) {
	randomBytes, err := utilities.GenRandomByteSlice(32)
	assert.NoError(t, err)
	t.Logf("Generated random byte: %s", string(randomBytes))
}
