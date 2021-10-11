package faker

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestFakeKeyName(t *testing.T) {
	name := Key().Name()
	namePrefix := strings.Split(name, "-")[0]
	assert.Contains(t, names, namePrefix, namePrefix)
}
