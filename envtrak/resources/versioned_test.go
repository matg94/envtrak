package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVVersionGetInt(t *testing.T) {
    
    version := VVersion{
        Value: "v11",
    }
    expectedNumericalVersion := 11

    assert.Equal(
        t,
        expectedNumericalVersion,
        version.GetInt(),
        "Expected versions to be the same",
    )

}

