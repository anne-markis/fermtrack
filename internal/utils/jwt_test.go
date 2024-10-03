package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateJWT(t *testing.T) {
	t.Run("generate and validate", func(t *testing.T) {
		os.Setenv("JWT_SECRET_KEY", "cooltest")
		defer func() {
			os.Setenv("JWT_SECRET_KEY", "")
		}()

		token, err := GenerateJWT("username")
		require.Nil(t, err)
		assert.NotEmpty(t, token)

		validToken, err := ValidateJWT(token)
		require.Nil(t, err)
		require.NotNil(t, validToken)

	})
}
