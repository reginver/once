package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameFromImageRef(t *testing.T) {
	assert.Equal(t, "once-campfire", NameFromImageRef("ghcr.io/basecamp/once-campfire:main"))
	assert.Equal(t, "once-campfire", NameFromImageRef("ghcr.io/basecamp/once-campfire"))
	assert.Equal(t, "nginx", NameFromImageRef("nginx:latest"))
	assert.Equal(t, "nginx", NameFromImageRef("nginx"))
}
