package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupVaultTokenFromFile(t *testing.T) {
	token, err := lookupVaultTokenFromFile("file does not exist")
	assert.Empty(t, token)
	assert.Error(t, err)

	token, err = lookupVaultTokenFromFile("internal/config/testdata/init.out.bad")
	assert.Empty(t, token)
	assert.Error(t, err)

	token, err = lookupVaultTokenFromFile("internal/config/testdata/init.out.good")
	assert.NoError(t, err)
	assert.Equal(t, "hvs.xAIi0RxVOTfwSNYivBhb3Gfp", token)
}
